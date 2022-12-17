package grpcext

import (
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	"qeco.dev/pkg/errs/internal"
)

type grpcError internal.CanonicalError

func (e *grpcError) GRPCStatus() *grpcstatus.Status {
	ce := (*internal.CanonicalError)(e)
	return grpcstatus.New(codes.Code(ce.Code()), ce.Error())
}

func (e *grpcError) Error() string {
	ce := (*internal.CanonicalError)(e)
	return ce.Error()
}

func (e *grpcError) Unwrap() error {
	ce := (*internal.CanonicalError)(e)
	return ce.Unwrap()
}

func GRPCErrorAdapter(err error) error {
	if err == nil {
		return nil
	}

	switch v := err.(type) {
	case *internal.CanonicalError:
		return (*grpcError)(v)
	default:
		return err
	}
}
