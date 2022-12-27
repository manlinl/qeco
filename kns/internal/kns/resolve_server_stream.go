package kns

import (
	"google.golang.org/grpc/status"

	pb "qeco.dev/apis/kns/v1"
)

type ResolveServerStream struct {
	id       int64
	resolver *Resolver
	stream   pb.NameService_StreamingResolveServer
	reqCh    chan *pb.StreamingResolveRequest
	errCh    chan error
}

func NewResolveServerStream(id int64, resolver *Resolver,
	stream pb.NameService_StreamingResolveServer) *ResolveServerStream {
	return &ResolveServerStream{
		id:       id,
		resolver: resolver,
		stream:   stream,
		reqCh:    make(chan *pb.StreamingResolveRequest),
		errCh:    make(chan error, 1),
	}
}

func (s *ResolveServerStream) Process() error {
	notifyCh := s.resolver.Register(s.id)
	defer s.resolver.Deregister(s.id)

	// Start to receive StreamingResolveRequest on a background Goroutine.
	go s.receiveRequests()
	strCtx := s.stream.Context()
	for {
		select {
		case <-strCtx.Done():
			return status.FromContextError(strCtx.Err()).Err()
		case req := <-s.reqCh:
			resp, err := s.generateResponse(req)
			if err != nil {
				return err
			}
			if err := s.stream.Send(resp); err != nil {
				return err
			}
		case resp := <-notifyCh:
			if err := s.stream.Send(resp); err != nil {
				return err
			}
		}
	}
}

func (s *ResolveServerStream) receiveRequests() {
	for {
		req, err := s.stream.Recv()
		if err != nil {
			s.errCh <- err
			return
		}
		s.reqCh <- req
	}
}

func (s *ResolveServerStream) generateResponse(
	req *pb.StreamingResolveRequest) (*pb.StreamingResolveResponse, error) {
	updates, err := s.resolver.UpdateSubscription(s.id, req.GetNamesSubscribe(),
		req.GetNamesUnsubscribes(), req.Option)
	if err != nil {
		return nil, err
	}
	return &pb.StreamingResolveResponse{
		Option:  req.Option,
		Updates: updates,
	}, nil
}
