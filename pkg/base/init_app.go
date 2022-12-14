package base

import (
	"flag"
	"k8s.io/klog/v2"
)

func InitApp() (cleanupFn func()) {
	klog.InitFlags(nil)
	flag.Parse()
	return func() {
		klog.Flush()
	}
}
