package pkg

import (
	"time"

	pb "qeco.dev/apis/kns/v1"
)

type (
	KNSBackend interface {
		Register(name, address string, ttl time.Duration) error
		ChangeEventChannel() chan ChangeEvent
		Resolve(name string) (*pb.ResolutionResult, error)
		ResolveMultiple(names []string) ([]*pb.ResolutionResult, error)
	}

	ChangeEvent struct {
		UpdatedNames []string
		RemovedNames []string
	}
)
