package dnsproxy

import (
	"context"
	"fmt"
	"net"

	"github.com/miekg/dns"
	"k8s.io/klog/v2"

	"qeco.dev/pkg/base"
)

type Server struct {
	domain   string
	dnsServ  *dns.Server
	resolver *KNSResolver
}

func NewServer(url string, domain string, resolver *KNSResolver) *Server {
	net, addr := base.MustParseAddress(url)
	mux := dns.NewServeMux()
	srv := &Server{
		domain: dns.CanonicalName(domain),
		dnsServ: &dns.Server{
			Addr:    addr,
			Net:     net,
			Handler: mux,
		},

		resolver: resolver,
	}
	mux.HandleFunc(".", srv.handleRequest)
	return srv
}

func (s *Server) Run() {
	if err := s.dnsServ.ListenAndServe(); err != nil {
		klog.Fatalf("Failed to serve DNS requests: %v", err)
	}
}

func (s *Server) Shutdown() {
	_ = s.dnsServ.Shutdown()
}

func (s *Server) handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	if !s.validate(w, r) {
		return
	}

	if !dns.IsSubDomain(s.domain, r.Question[0].Name) {
		s.queryUpstreams(w, r)
		return
	}

	s.handleKNSRequest(w, r)
}

func (s *Server) validate(w dns.ResponseWriter, r *dns.Msg) bool {
	if len(r.Question) != 1 ||
		r.Question[0].Qtype != dns.TypeA ||
		r.Question[0].Qclass != dns.ClassINET {
		reply := new(dns.Msg)
		reply.SetReply(r)
		reply.Rcode = dns.RcodeFormatError
		_ = w.WriteMsg(reply)
		return false
	}
	return true
}

func (s *Server) handleKNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	if !s.sanityCheckKNSRequest(w, r) {
		return
	}

	reply := new(dns.Msg)
	reply.SetReply(r)
	reply.Authoritative = true
	defer w.WriteMsg(reply)
	name := r.Question[0].Name
	result, err := s.resolver.Resolve(context.Background(), name)
	if err != nil {
		reply.Rcode = dns.RcodeServerFailure
		return
	}

	if result.GetStatus() != nil {
		reply.Rcode = dns.RcodeNameError
		return
	}

	for _, ip := range result.GetAddresses() {
		if rr, err := dns.NewRR(fmt.Sprintf("%s A %s", name, ip)); err == nil {
			reply.Answer = append(reply.Answer, rr)
		}
	}
}

func (s *Server) sanityCheckKNSRequest(w dns.ResponseWriter, r *dns.Msg) bool {
	if r.Opcode != dns.OpcodeQuery ||
		r.Question[0].Qtype != dns.TypeA ||
		r.Question[0].Qclass != dns.ClassINET {
		reply := new(dns.Msg)
		reply.SetReply(r)
		reply.Rcode = dns.RcodeServerFailure
		_ = w.WriteMsg(reply)
	}
	return true
}

func (s *Server) queryUpstreams(w dns.ResponseWriter, r *dns.Msg) {
	reply := new(dns.Msg)
	reply.SetReply(r)
	defer w.WriteMsg(reply)
	host := r.Question[0].Name
	ips, err := net.DefaultResolver.LookupIPAddr(context.Background(), host)
	if err != nil {
		reply.Rcode = dns.RcodeServerFailure
		return
	}

	for _, ip := range ips {
		var rtype = "AAAA"
		if p4 := ip.IP.To4(); len(p4) == net.IPv4len {
			rtype = "A"
		}

		rr, err := dns.NewRR(fmt.Sprintf("%s %d %s %s", host, 10, rtype, ip.String()))
		if err != nil {
			reply.Rcode = dns.RcodeServerFailure
			return
		}
		reply.Answer = append(reply.Answer, rr)
	}
}
