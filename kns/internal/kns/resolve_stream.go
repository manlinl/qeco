package kns

import (
	"google.golang.org/grpc/status"

	pb "qeco.dev/apis/kns/v1"
)

type ResolveStream struct {
	id       int64
	resolver *Resolver
	stream   pb.NameService_StreamingResolveServer
	reqCh    chan *pb.StreamingResolveRequest
	errCh    chan error
}

func NewResolveStream(id int64, resolver *Resolver,
	stream pb.NameService_StreamingResolveServer) *ResolveStream {
	return &ResolveStream{
		id:       id,
		resolver: resolver,
		stream:   stream,
		reqCh:    make(chan *pb.StreamingResolveRequest),
		errCh:    make(chan error, 1),
	}
}

func (r *ResolveStream) Process() error {
	notifyCh := r.resolver.Register(r.id)
	defer r.resolver.Deregister(r.id)

	// Start to receive StreamingResolveRequest on a background Goroutine.
	go r.receiveRequests()
	strCtx := r.stream.Context()
	for {
		select {
		case <-strCtx.Done():
			return status.FromContextError(strCtx.Err()).Err()
		case req := <-r.reqCh:
			resp, err := r.generateResponse(req)
			if err != nil {
				return err
			}
			if err := r.stream.Send(resp); err != nil {
				return err
			}
		case resp := <-notifyCh:
			if err := r.stream.Send(resp); err != nil {
				return err
			}
		}
	}
}

func (r *ResolveStream) receiveRequests() {
	for {
		req, err := r.stream.Recv()
		if err != nil {
			r.errCh <- err
			return
		}
		r.reqCh <- req
	}
}

func (r *ResolveStream) generateResponse(
	req *pb.StreamingResolveRequest) (*pb.StreamingResolveResponse, error) {
	updates, err := r.resolver.UpdateSubscription(r.id, req.GetNamesSubscribe(),
		req.GetNamesUnsubscribes(), req.Option)
	if err != nil {
		return nil, err
	}
	return &pb.StreamingResolveResponse{
		Option:  req.Option,
		Updates: updates,
	}, nil
}
