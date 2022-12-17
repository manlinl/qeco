package grpcext

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	"qeco.dev/pkg/errs"
)

func TestGRPCErrorAdapter(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		assert.Nil(t, GRPCErrorAdapter(nil))
	})

	t.Run("CanonicalError", func(t *testing.T) {
		err := GRPCErrorAdapter(errs.Error(errs.NotFound, "NotFound"))
		assert.Equal(t, codes.NotFound, grpcstatus.Convert(err).Code())
		assert.Equal(t, "NotFound", grpcstatus.Convert(err).Message())
		assert.Equal(t, codes.NotFound, grpcstatus.Code(err))
	})
}
