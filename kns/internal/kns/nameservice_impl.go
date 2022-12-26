package kns

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/google/uuid"

	pb "qeco.dev/apis/kns/v1"
	"qeco.dev/kns/pkg"
	"qeco.dev/pkg/errs"
	"qeco.dev/pkg/errs/grpcext"
)

type NameServiceImpl struct {
	pb.UnimplementedNameServiceServer
	ttl             time.Duration
	backend         pkg.KNSBackend
	resolver        *Resolver
	resolveStreamId int64
}

func NewNameServiceImpl(ttl time.Duration, backend pkg.KNSBackend,
	option ResolverOption) *NameServiceImpl {
	return &NameServiceImpl{
		ttl:             ttl,
		backend:         backend,
		resolver:        NewResolver(backend, option),
		resolveStreamId: 0,
	}
}

func (s *NameServiceImpl) Register(stream pb.NameService_RegisterServer) error {
	err := NewRegisterStream(uuid.New().String(), stream, s.ttl, s.backend).Process()
	return grpcext.GRPCErrorAdapter(err)
}

func (s *NameServiceImpl) Resolve(ctx context.Context,
	request *pb.ResolveRequest) (*pb.ResolveResponse, error) {
	name := request.GetName()
	if len(name) == 0 {
		return nil, errs.Error(errs.InvalidArgument,
			"ResolveRequest must not have empty name field")
	}

	result, err := s.backend.Resolve(name)
	if err != nil {
		return nil, err
	}
	return &pb.ResolveResponse{
		Result: result,
	}, nil
}

func (s *NameServiceImpl) StreamingResolve(stream pb.NameService_StreamingResolveServer) error {
	err := NewResolveStream(s.nextResolveStreamID(), s.resolver, stream).Process()
	return grpcext.GRPCErrorAdapter(err)
}

func (s *NameServiceImpl) nextResolveStreamID() int64 {
	return atomic.AddInt64(&s.resolveStreamId, 1)
}
