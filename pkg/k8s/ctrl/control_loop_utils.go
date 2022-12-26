package ctrl

import "k8s.io/client-go/util/workqueue"

func HandleErr(queue workqueue.RateLimitingInterface, err error, retries int, key interface{}) {
	if err == nil {
		queue.Forget(key)
		return
	}

	if queue.NumRequeues(key) < retries {
		queue.AddRateLimited(key)
		return
	}

	queue.Forget(key)
}
