package base

import (
	"sync"
)

type GracefulShutdown struct {
	wg sync.WaitGroup
}

func (l *GracefulShutdown) Run(calls ...func()) {
	for _, c := range calls {
		go func(c func()) {
			l.wg.Add(1)
			defer l.wg.Done()
			c()
		}(c)
	}
}

func (l *GracefulShutdown) Wait() {
	l.wg.Wait()
}
