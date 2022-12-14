package pkg

import "time"

type (
	KVStore interface {
		Set(kv KV, ttl time.Duration) error
		Get(key string) (value string, err error)
		Del(key string) bool
		Watch(prefix string) (<-chan KV, CancelFn, error)
	}

	CancelFn = func()

	KV struct {
		Key   string
		Value string
	}
)
