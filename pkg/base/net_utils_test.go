package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustParseAddress(t *testing.T) {
	t.Run("TCP", func(t *testing.T) {
		net, addr := MustParseAddress("tcp://localhost:53")
		assert.Equal(t, "tcp", net)
		assert.Equal(t, "localhost:53", addr)
	})
	t.Run("TCP/PortOnly", func(t *testing.T) {
		net, addr := MustParseAddress("tcp://:53")
		assert.Equal(t, "tcp", net)
		assert.Equal(t, ":53", addr)
	})
	t.Run("UDP", func(t *testing.T) {
		net, addr := MustParseAddress("udp://localhost:53")
		assert.Equal(t, "udp", net)
		assert.Equal(t, "localhost:53", addr)
	})
	t.Run("UDP/PortOnly", func(t *testing.T) {
		net, addr := MustParseAddress("udp://:53")
		assert.Equal(t, "udp", net)
		assert.Equal(t, ":53", addr)
	})
	t.Run("UDS/PortOnly", func(t *testing.T) {
		net, addr := MustParseAddress("unix:///tmp/zzz.sock")
		assert.Equal(t, "unix", net)
		assert.Equal(t, "/tmp/zzz.sock", addr)
	})
}
