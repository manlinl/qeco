package client

import (
	"net/netip"
	"time"

	"k8s.io/klog/v2"
	pb "qeco.dev/apis/kns/v1"
)

type RegisterClientStream struct {
	knsName string
	ip      netip.Addr
	stream  pb.NameService_RegisterClient
	respCh  chan *pb.RegisterResponse
	errCh   chan error
}

func NewRegisterClientStream(stream pb.NameService_RegisterClient,
	knsName string, ip netip.Addr) *RegisterClientStream {
	return &RegisterClientStream{
		knsName: knsName,
		ip:      ip,
		stream:  stream,
		respCh:  make(chan *pb.RegisterResponse),
		errCh:   make(chan error, 1),
	}
}

func (s *RegisterClientStream) Process() (err error) {
	klog.V(4).InfoS("Start Register stream", "stream", s.stream)
	defer klog.V(4).InfoS("Finish Register stream", "stream", s.stream, "error", err)

	go s.receiveRegisterResponse()

	if err = s.sendRegisterRequest(); err != nil {
		return err
	}

	timer := time.NewTimer(1 * time.Hour)
	for {
		select {
		case resp := <-s.respCh:
			timer.Reset(computeTimerDuration(resp.GetTtl()))
		case err = <-s.errCh:
			return
		case <-timer.C:
			if err = s.sendRegisterRequest(); err != nil {
				return
			}
		}
	}
}

func (s *RegisterClientStream) receiveRegisterResponse() {
	klog.V(4).InfoS("Start receiving RegisterResponse", "stream", s.stream)
	defer klog.V(4).InfoS("End receiving RegisterResponse", "stream", s.stream)

	for {
		resp, err := s.stream.Recv()
		klog.V(4).InfoS("Receive RegisterResponse", "stream", s.stream, "error", err)
		if err != nil {
			s.errCh <- err
			return
		}
		s.respCh <- resp
	}
}

func (s *RegisterClientStream) sendRegisterRequest() error {
	request := pb.RegisterRequest{
		Name:    s.knsName,
		Address: s.ip.String(),
	}
	err := s.stream.Send(&request)
	klog.V(4).InfoS("Send RegisterRequest", "stream", s.stream, "error", err)
	return err
}
