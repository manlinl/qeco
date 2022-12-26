package mem

import (
	"context"
	"fmt"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "qeco.dev/apis/kns/v1"
	"qeco.dev/kns/pkg"
	"qeco.dev/pkg/errs"
)

type (
	InMemBackend struct {
		cache         *ttlcache.Cache[string, string]
		changeEventCh chan pkg.ChangeEvent
	}
)

func NewInMemBackend() *InMemBackend {
	changeEventCh := make(chan pkg.ChangeEvent, 100)
	cache := ttlcache.New[string, string]()
	cache.OnInsertion(func(ctx context.Context, item *ttlcache.Item[string, string]) {
		changeEventCh <- pkg.ChangeEvent{
			UpdatedNames: []string{item.Key()},
		}
	})
	cache.OnEviction(func(ctx context.Context, reason ttlcache.EvictionReason,
		item *ttlcache.Item[string, string]) {
		changeEventCh <- pkg.ChangeEvent{
			RemovedNames: []string{item.Key()},
		}
	})

	return &InMemBackend{
		cache:         cache,
		changeEventCh: changeEventCh,
	}
}

func (b *InMemBackend) Register(name, address string, ttl time.Duration) error {
	b.cache.Set(name, address, ttl)
	return nil
}

func (b *InMemBackend) ChangeEventChannel() chan pkg.ChangeEvent {
	return b.changeEventCh
}

func (b *InMemBackend) Resolve(name string) (*pb.ResolutionResult, error) {
	record := b.cache.Get(name)
	if record == nil {
		return &pb.ResolutionResult{
			Name: name,
			Status: &status.Status{
				Code:    int32(errs.NotFound),
				Message: fmt.Sprintf("Name is not found: %s", name),
			},
		}, nil
	}
	return &pb.ResolutionResult{
		Name:      name,
		Addresses: nil,
		Ttl:       timestamppb.New(record.ExpiresAt()),
	}, nil
}

func (b *InMemBackend) ResolveMultiple(names []string) ([]*pb.ResolutionResult, error) {
	var results []*pb.ResolutionResult
	for _, name := range names {
		result, err := b.Resolve(name)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
