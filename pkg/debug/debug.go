package debug

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/klog/v2"
)

var debugPort = flag.Int("debug_port", 9090, "Port exposing various debug information.")

func init() {
	http.Handle("/debug/metrics", promhttp.Handler())
	http.Handle("/metrics", promhttp.Handler())
}

func StartDebugPage() {
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", *debugPort), nil)
		if err != nil && err != http.ErrServerClosed {
			klog.Fatalf("Failed to tart zPage service.")
		}
	}()
}
