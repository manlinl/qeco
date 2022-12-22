package kns

import (
	"context"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/durationpb"

	pb "qeco.dev/apis/kns/v1"
	"qeco.dev/kns/pkg"
	"qeco.dev/pkg/errs"
	"qeco.dev/pkg/errs/grpcext"
)

type NameServiceImpl struct {
	pb.UnimplementedNameServiceServer
	ttl     time.Duration
	storage pkg.KVStore
}

func NewNameServiceImpl(ttl time.Duration, storage pkg.KVStore) *NameServiceImpl {
	return &NameServiceImpl{
		ttl:     ttl,
		storage: storage,
	}
}

func (s *NameServiceImpl) Register(stream pb.NameService_RegisterServer) error {
	err := NewRegisterStream(uuid.New().String(), stream, s.ttl, s.storage).Process()
	return grpcext.GRPCErrorAdapter(err)
}

func (s *NameServiceImpl) Resolve(ctx context.Context,
	request *pb.ResolveRequest) (*pb.ResolveResponse, error) {
	name := request.GetName()
	if len(name) == 0 {
		return nil, errs.Error(errs.InvalidArgument,
			"ResolveRequest must not have empty name field")
	}

	value, err := s.storage.Get(name)
	if err != nil {
		return nil, err
	}
	return &pb.ResolveResponse{
		Name:      name,
		Addresses: []string{value},
		Ttl:       durationpb.New(s.ttl),
	}, nil
}
