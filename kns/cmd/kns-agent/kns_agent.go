package main

import (
	"context"
	"flag"
	"net/netip"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"k8s.io/klog/v2"

	pb "qeco.dev/apis/kns/v1"
	"qeco.dev/kns/internal/client"
	"qeco.dev/kns/internal/dnsproxy"
	"qeco.dev/pkg/base"
	_ "qeco.dev/pkg/debug"
)

var (
	knsServer      = flag.String("kns_server", ":9290", "KNS server address.")
	dnsName        = flag.String("dns_name", "localhost.kns.qeco.dev", "DNS name registered to KNS server.")
	ip             = flag.String("ip", "127.0.0.1", "IP to register with KNS server.")
	dnsServer      = flag.String("dns_server", "udp://:5300", "DNS server endpoint.")
	minTTL         = flag.Duration("min_ttl", 1*time.Second, "Minimum TTL for DNS record.")
	localCacheSize = flag.Int("name_cache_size", 1000, "Local name cache size.")
	knsDomain      = flag.String("kns_domain", "kns.qeco.dev", "KNS root domain")
)

func main() {
	cleanupFn := base.InitApp()
	defer cleanupFn()

	klog.InfoS("Run kns-agent")

	conn, err := grpc.Dial(*knsServer,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		klog.Fatalf("Fail to create connection to KNS Server %v", err)
	}

	stopCh := base.SetupShutdownSignalHandler()
	knsClient := pb.NewNameServiceClient(conn)

	registrar := getRegistrar(knsClient)
	dnsProxy := getDNSProxy(knsClient)

	ctx, cancelFn := context.WithCancel(context.Background())
	var launcher base.GracefulShutdown
	launcher.Run(
		func() { registrar.Run(ctx) },
		func() { dnsProxy.Run() },
	)
	defer launcher.Wait()
	defer cancelFn()
	defer dnsProxy.Shutdown()
	<-stopCh
}

func getRegistrar(knsClient pb.NameServiceClient) *client.Registrar {
	ipAddr := netip.MustParseAddr(*ip)
	return client.NewRegistrar(*dnsName, ipAddr, knsClient)
}

func getDNSProxy(knsClient pb.NameServiceClient) *dnsproxy.Server {
	resolver := dnsproxy.NewKNSResolver(knsClient, dnsproxy.ResolverOption{
		MinTTL:        *minTTL,
		NameCacheSize: *localCacheSize,
	})
	return dnsproxy.NewServer(*dnsServer, *knsDomain, resolver)
}
