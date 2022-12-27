package base

import (
	"log"
	"net/url"
)

func MustParseAddress(address string) (network string, addr string) {
	url, err := url.Parse(address)
	if err != nil {
		log.Fatalf("Failed to parse adress: %q", address)
	}
	switch url.Scheme {
	case "tcp", "udp":
		return url.Scheme, url.Host
	case "unix":
		return "unix", url.Path
	default:
		log.Fatalf("Unspported address type: %s", address)
		return
	}
}
