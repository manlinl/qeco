package kns

import (
	"sync"
	"time"

	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	pb "qeco.dev/apis/kns/v1"
	"qeco.dev/kns/pkg"
	"qeco.dev/pkg/errs"
	"qeco.dev/pkg/k8s/ctrl"
)

type (
	ResolverOption struct {
		ListenerWorkerCount            int
		NotificationWorkerCount        int
		NotificationChannelBufferSize  int
		NotificationChannelSendTimeout time.Duration
	}

	Resolver struct {
		backend pkg.KNSBackend
		option  ResolverOption
		queue   workqueue.RateLimitingInterface

		// mtx guards `streams` and `index` fields below it.
		mtx     sync.RWMutex
		streams map[int64]*streamInfo
		index   map[string]streamInfoMap
	}

	streamInfo struct {
		id         int64
		notifyCh   chan *pb.StreamingResolveResponse
		subscribed map[string]bool
	}

	// Convenient type aliases.
	streamInfoMap = map[int64]*streamInfo
)

func NewResolver(backend pkg.KNSBackend, option ResolverOption) *Resolver {
	return &Resolver{
		backend: backend,
		option:  option,
		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(),
			"NameResolver"),
		streams: make(map[int64]*streamInfo),
		index:   make(map[string]streamInfoMap),
	}
}

func (r *Resolver) Run(stopCh <-chan struct{}) {
	for i := 0; i < r.option.ListenerWorkerCount; i++ {
		go r.startListenerWorker(stopCh)
	}

	defer r.queue.ShutDown()
	for i := 0; i < r.option.NotificationWorkerCount; i++ {
		go r.startNotificationWorker()
	}
	<-stopCh
}

func (r *Resolver) Register(streamId int64) chan *pb.StreamingResolveResponse {
	strInfo := &streamInfo{
		id:         streamId,
		notifyCh:   make(chan *pb.StreamingResolveResponse, r.option.NotificationChannelBufferSize),
		subscribed: make(map[string]bool),
	}
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, found := r.streams[streamId]; found {
		klog.Fatalf("Stream has been registered, id: %d", streamId)
	}
	r.streams[streamId] = strInfo
	return strInfo.notifyCh
}

func (r *Resolver) Deregister(streamId int64) {
	strInfo, found := r.getStreamInfo(streamId, true)
	if !found {
		klog.ErrorS(
			errs.Error(errs.Internal, "Stream is not registered"),
			"Stream is not registered", "id", streamId)
		return
	}

	func() {
		r.mtx.Lock()
		defer r.mtx.Unlock()
		for name := range strInfo.subscribed {
			if infoMap, found := r.index[name]; found {
				delete(infoMap, streamId)
			}
		}
	}()
}

func (r *Resolver) UpdateSubscription(
	streamId int64, namesSubscribe, namesUnsubscribe []string,
	option pb.StreamResolveOption) ([]*pb.ResolutionResult, error) {
	strInfo, found := r.getStreamInfo(streamId, false)
	if !found {
		return nil, errs.Errorf(errs.Internal, "Stream is not registered, id: %d", streamId)
	}

	for _, name := range namesSubscribe {
		strInfo.subscribed[name] = true
	}

	for _, unsubscribed := range namesUnsubscribe {
		delete(strInfo.subscribed, unsubscribed)
	}

	func() {
		r.mtx.Lock()
		defer r.mtx.Unlock()
		for name := range strInfo.subscribed {
			r.index[name][streamId] = strInfo
		}

		for _, unsubscribed := range namesUnsubscribe {
			delete(r.index[unsubscribed], streamId)
		}
	}()

	switch option {
	case pb.StreamResolveOption_STREAM_RESOLVE_OPTION_DELTA:
		return r.backend.ResolveMultiple(namesSubscribe)
	case pb.StreamResolveOption_STREAM_RESOLVE_OPTION_ALL:
		return r.backend.ResolveMultiple(toStringSlice(strInfo.subscribed))
	default:
		return nil, errs.Errorf(errs.InvalidArgument, "Unknown StreamResolveOption: %v", option)
	}
}

func (r *Resolver) startListenerWorker(stopCh <-chan struct{}) {
	notificationCh := r.backend.ChangeEventChannel()
	for {
		select {
		case <-stopCh:
			return
		case changeEvent := <-notificationCh:
			for updatedName := range changeEvent.UpdatedNames {
				r.queue.Add(updatedName)
			}

			for deletedName := range changeEvent.RemovedNames {
				r.queue.Add(deletedName)
			}
		}
	}
}

func (r *Resolver) startNotificationWorker() {
	for r.processNextWorkItem() {
	}
}

func (r *Resolver) processNextWorkItem() bool {
	obj, shutdown := r.queue.Get()
	if shutdown {
		return false
	}
	defer r.queue.Done(obj)
	ctrl.HandleErr(r.queue, r.doResolve(obj.(string)), 5, obj)
	return true
}

func (r *Resolver) doResolve(name string) error {
	result, err := r.backend.Resolve(name)
	if err != nil {
		return err
	}

	resp := &pb.StreamingResolveResponse{
		Option:  pb.StreamResolveOption_STREAM_RESOLVE_OPTION_DELTA,
		Updates: []*pb.ResolutionResult{result},
	}

	infoList := r.getStreamInfoList(name)
	timer := time.NewTimer(r.option.NotificationChannelSendTimeout)
	defer timer.Stop()
	for _, strInfo := range infoList {
		select {
		case <-timer.C:
			klog.V(3).InfoS("Failed to send resole result", "streamId", strInfo.id)
		case strInfo.notifyCh <- resp:
		}
		timer.Reset(r.option.NotificationChannelSendTimeout)
	}
	return nil
}

func (r *Resolver) getStreamInfo(streamId int64, remove bool) (info *streamInfo, found bool) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	info, found = r.streams[streamId]
	if found && remove {
		delete(r.streams, streamId)
	}
	return
}

func (r *Resolver) getStreamInfoList(name string) []*streamInfo {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	infoMap := r.index[name]
	streamInfoList := make([]*streamInfo, len(infoMap))
	for _, info := range r.index[name] {
		streamInfoList = append(streamInfoList, info)
	}
	return streamInfoList
}

func toStringSlice(names map[string]bool) []string {
	if len(names) == 0 {
		return nil
	}

	slice := make([]string, len(names))
	for name := range names {
		slice = append(slice, name)
	}
	return slice
}
