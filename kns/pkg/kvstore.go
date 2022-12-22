package pkg

import "time"

const (
	ChangeTypeInvalid KVChangeType = iota
	ChangeTypeCreate
	ChangeTypeUpdate
	ChangeTypeDeleted
)

type (
	KVStore interface {
		Set(kv KV, ttl time.Duration) error
		Get(key string) (value string, err error)
		Del(key string)
		Watch(prefix string) (<-chan *KVChangeEvent, WatchCancelFn, error)
	}

	WatchCancelFn = func()

	KV struct {
		Key   string
		Value string
	}

	KVChangeType int

	KVChangeEvent struct {
		Type   KVChangeType
		KVPair KV
	}
)
