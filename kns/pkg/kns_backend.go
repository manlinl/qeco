package pkg

import (
	"time"

	pb "qeco.dev/apis/kns/v1"
)

type (
	KNSBackend interface {
		Register(name, address string, ttl time.Duration) error
		ChangeEventChannel() chan ChangeEvent
		Resolve(name string) (*pb.ResolveResult, error)
		ResolveMultiple(names []string) ([]*pb.ResolveResult, error)
	}

	ChangeEvent struct {
		UpdatedNames []string
		RemovedNames []string
	}
)
