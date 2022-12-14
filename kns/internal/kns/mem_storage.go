package kns

import (
	"time"

	"qeco.dev/kns/pkg"
)

type MemStorage struct{}

func (f *MemStorage) Set(kv pkg.KV, ttl time.Duration) error {
	return nil
}

func (f *MemStorage) Get(key string) (value string, err error) {
	return "", nil
}

func (f *MemStorage) Del(key string) bool {
	return false
}

func (f *MemStorage) Watch(prefix string) (<-chan pkg.KV, pkg.CancelFn, error) {
	return nil, nil, nil
}
