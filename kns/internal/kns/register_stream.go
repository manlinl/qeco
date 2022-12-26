package kns

import (
	"time"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	pb "qeco.dev/apis/kns/v1"
	"qeco.dev/kns/pkg"
)

type (
	RegisterStream struct {
		stream  pb.NameService_RegisterServer
		ttl     time.Duration
		backend pkg.KNSBackend
		id      string
		reqCh   chan *pb.RegisterRequest
		errCh   chan error
		record  *nameRecord
	}

	nameRecord struct {
		name    string
		address string
	}
)

func NewRegisterStream(id string, stream pb.NameService_RegisterServer,
	ttl time.Duration, backend pkg.KNSBackend) *RegisterStream {
	return &RegisterStream{
		id:      id,
		stream:  stream,
		ttl:     ttl,
		backend: backend,
		reqCh:   make(chan *pb.RegisterRequest),
		errCh:   make(chan error, 1),
		record:  nil,
	}
}

func (r *RegisterStream) Process() error {
	klog.V(3).InfoS("Start processing stream", "stream", r.id)
	defer func() {
		klog.V(3).InfoS("Finish processing stream", "stream", r.id)
	}()

	go r.receiveRequests()
	if err := r.handleFirstRequest(); err != nil {
		return err
	}

	strCtx := r.stream.Context()
	for {
		select {
		case <-strCtx.Done():
			klog.V(4).InfoS("Stream context done", "stream", r.id, "err", strCtx.Err())
			return status.FromContextError(strCtx.Err()).Err()
		case <-r.reqCh:
			klog.V(3).InfoS("Receive RegisterRequest", "stream", r.id)
			if err := r.sendRegisterResponse(); err != nil {
				return err
			}
		case err := <-r.errCh:
			return err
		}
	}
}

func (r *RegisterStream) receiveRequests() {
	klog.V(4).InfoS("Start receiving RegisterRequest", "stream", r.id)
	defer klog.V(4).InfoS("Finish receiving RegisterRequest", "stream", r.id)

	for {
		req, err := r.stream.Recv()
		if err != nil {
			klog.V(2).ErrorS(err, "Error on receiving RegisterRequest", "stream", r.id)
			r.errCh <- err
			return
		} else {
			r.reqCh <- req
		}
	}
}

func (r *RegisterStream) handleFirstRequest() error {
	select {
	case req := <-r.reqCh:
		r.record = &nameRecord{
			name:    req.GetName(),
			address: req.GetAddress(),
		}
		return r.sendRegisterResponse()
	case err := <-r.errCh:
		return err
	}
}

func (r *RegisterStream) sendRegisterResponse() (err error) {
	if err = r.backend.Register(r.record.name, r.record.address, r.ttl); err != nil {
		return
	}
	resp := pb.RegisterResponse{
		Id:  r.id,
		Ttl: timestamppb.New(time.Now().Add(r.ttl)),
	}
	if err = r.stream.Send(&resp); err != nil {
		return
	}
	return nil
}
