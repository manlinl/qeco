package kns

import (
	"time"

	"github.com/google/uuid"

	pb "qeco.dev/apis/kns/v1"
	"qeco.dev/kns/pkg"
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
	return NewRegisterStream(uuid.New().String(), stream, s.ttl, s.storage).Process()
}
