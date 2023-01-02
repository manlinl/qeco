package dnsproxy

import (
	"context"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/jonboulle/clockwork"

	pb "qeco.dev/apis/kns/v1"
)

type (
	KNSResolver struct {
		knsClient pb.NameServiceClient
		option    ResolverOption
		clock     clockwork.Clock
		cache     *ttlcache.Cache[string, *pb.ResolveResult]
	}

	ResolverOption struct {
		MinTTL        time.Duration
		NameCacheSize int
	}
)

func NewKNSResolver(knsClient pb.NameServiceClient, option ResolverOption) *KNSResolver {
	cache := ttlcache.New[string, *pb.ResolveResult](
		ttlcache.WithCapacity[string, *pb.ResolveResult](uint64(option.NameCacheSize)),
		ttlcache.WithDisableTouchOnHit[string, *pb.ResolveResult]())

	return &KNSResolver{
		knsClient: knsClient,
		option:    option,
		clock:     clockwork.NewRealClock(),
		cache:     cache,
	}
}

func (r *KNSResolver) Resolve(ctx context.Context, name string) (*pb.ResolveResult, error) {
	item := r.cache.Get(name, ttlcache.WithDisableTouchOnHit[string, *pb.ResolveResult]())
	if item != nil {
		return item.Value(), nil
	}

	resp, err := r.knsClient.Resolve(ctx, &pb.ResolveRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	ttl := resp.GetResult().GetTtl().AsDuration()
	if ttl < r.option.MinTTL {
		ttl = r.option.MinTTL
	}
	r.cache.Set(name, resp.GetResult(), ttl)
	return resp.GetResult(), nil
}
