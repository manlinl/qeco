package client

import (
	"context"
	"io"
	"net/netip"
	"time"

	mdns "github.com/miekg/dns"
	"k8s.io/klog/v2"

	"github.com/cenkalti/backoff/v4"

	pb "qeco.dev/apis/kns/v1"
)

type Registrar struct {
	ksnName   string
	ip        netip.Addr
	knsClient pb.NameServiceClient
}

func NewRegistrar(name string, ip netip.Addr, knsClient pb.NameServiceClient) *Registrar {
	return &Registrar{
		ksnName:   mdns.CanonicalName(name),
		ip:        ip,
		knsClient: knsClient,
	}
}

func (r *Registrar) Run(ctx context.Context) error {
	eb := getExponentialBackOff()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err := r.registerImpl(ctx)
			if err != nil && shouldBackOff(err) {
				backOff := eb.NextBackOff()
				klog.V(3).InfoS("Register() backs off", "error", err, "backOff", backOff)
				time.Sleep(backOff)
			}
		}
	}
}

func (r *Registrar) registerImpl(ctx context.Context) error {
	stream, err := r.knsClient.Register(ctx)
	if err != nil {
		return err
	}

	return NewRegisterClientStream(stream, r.ksnName, r.ip).Process()
}

func getExponentialBackOff() *backoff.ExponentialBackOff {
	eb := &backoff.ExponentialBackOff{
		InitialInterval:     backoff.DefaultInitialInterval,
		RandomizationFactor: backoff.DefaultRandomizationFactor,
		Multiplier:          backoff.DefaultMultiplier,
		MaxInterval:         5 * time.Second,
		MaxElapsedTime:      0, // Never stop
		Stop:                -1,
		Clock:               backoff.SystemClock,
	}
	eb.Reset()
	return eb
}

func shouldBackOff(err error) bool {
	return err != io.EOF
}
