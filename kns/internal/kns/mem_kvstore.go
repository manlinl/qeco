package kns

import (
	"sync"
	"time"

	"qeco.dev/pkg/errs"

	"qeco.dev/kns/pkg"
)

type (
	MemKVStore struct {
		sync.RWMutex
		kvs map[string]valueWithExpiry
	}

	valueWithExpiry struct {
		value  string
		expiry time.Time
	}
)

func NewMemKVStore() *MemKVStore {
	return &MemKVStore{
		RWMutex: sync.RWMutex{},
		kvs:     make(map[string]valueWithExpiry),
	}
}

func (f *MemKVStore) Set(kv pkg.KV, ttl time.Duration) error {
	f.Lock()
	defer f.Unlock()
	f.kvs[kv.Key] = valueWithExpiry{
		value:  kv.Value,
		expiry: time.Now().Add(ttl),
	}
	return nil
}

func (f *MemKVStore) Get(key string) (value string, err error) {
	f.Lock()
	defer f.Unlock()

	v, found := f.kvs[key]
	if !found {
		return "", errs.Errorf(errs.NotFound, "key %q is not found", key)
	} else if v.expiry.Before(time.Now()) {
		delete(f.kvs, key)
		return "", errs.Errorf(errs.NotFound, "key %q is not found", key)
	}
	return v.value, nil
}

func (f *MemKVStore) Del(key string) {
	f.Lock()
	defer f.Unlock()
	delete(f.kvs, key)
}

func (f *MemKVStore) Watch(prefix string) (<-chan *pkg.KVChangeEvent, pkg.WatchCancelFn, error) {
	return nil, nil, nil
}
