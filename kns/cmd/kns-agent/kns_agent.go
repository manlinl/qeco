package main

import (
	"context"
	"flag"
	"net/netip"

	"qeco.dev/pkg/base"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/klog/v2"
	pb "qeco.dev/apis/kns/v1"
	"qeco.dev/kns/internal/client"
	_ "qeco.dev/pkg/debug"
)

var (
	knsServer = flag.String("kns_server", "localhost:9290", "KNS server address.")
	ip        = flag.String("ip", "127.0.0.1", "IP to register with KNS server")
	address   = flag.String("address", "localhost", "Address used to register with KNS server")
)

func main() {
	cleanupFn := base.InitApp()
	defer cleanupFn()

	klog.InfoS("Start kns-agent")

	conn, err := grpc.Dial(*knsServer,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		klog.Fatalf("Fail to create connection to KNS Server %v", err)
	}

	ipAddr := netip.MustParseAddr(*ip)
	knsClient := pb.NewNameServiceClient(conn)
	registrar := client.NewRegistrar(*address, ipAddr, knsClient)
	_ = registrar.Run(context.Background())
}
