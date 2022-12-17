package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-prometheus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"qeco.dev/pkg/base"

	"k8s.io/klog/v2"

	pb "qeco.dev/apis/kns/v1"
	"qeco.dev/kns/internal/kns"
	"qeco.dev/pkg/debug"
	_ "qeco.dev/pkg/debug"
)

const (
	grpcKeepaliveTime        = 30 * time.Second
	grpcKeepaliveTimeout     = 10 * time.Second
	grpcKeepaliveMinTime     = 30 * time.Second
	grpcMaxConcurrentStreams = 1000000
)

var (
	port = flag.Int("port", 9290, "Kns server gRPC service port.")
	ttl  = flag.Duration("heart_beat_ttl", 10*time.Second, "Heart beat TTL for Register method.")
)

func main() {
	cleanupFn := base.InitApp()
	defer cleanupFn()
	debug.StartDebugPage()
	srv := startKNSServer()
	defer srv.GracefulStop()
	stopCh := base.SetupShutdownSignalHandler()
	<-stopCh
	klog.InfoS("Shut down kns-server")
}

func startKNSServer() *grpc.Server {
	grpcOptions := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    grpcKeepaliveTime,
			Timeout: grpcKeepaliveTimeout,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             grpcKeepaliveMinTime,
			PermitWithoutStream: true,
		}),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		)),
	}

	srv := grpc.NewServer(grpcOptions...)
	reflection.Register(srv)
	pb.RegisterNameServiceServer(srv, kns.NewNameServiceImpl(*ttl, &kns.MemKVStore{}))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		klog.Fatalf("Fail to create listener for KNS server %v", err)
	}

	klog.InfoS("Start KNS server", "addr", lis.Addr())
	go func() {
		if err := srv.Serve(lis); err != nil {
			klog.Fatalf("fail to start load balancer service: %v", err)
		}
	}()
	return srv
}
