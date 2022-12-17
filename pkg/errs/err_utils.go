package errs

import (
	"errors"
	"fmt"

	rpccode "google.golang.org/genproto/googleapis/rpc/code"

	"qeco.dev/pkg/errs/internal"
)

const (
	Ok                 = rpccode.Code_OK
	Cancelled          = rpccode.Code_CANCELLED
	Unknown            = rpccode.Code_UNKNOWN
	InvalidArgument    = rpccode.Code_INVALID_ARGUMENT
	DeadlineExceed     = rpccode.Code_DEADLINE_EXCEEDED
	NotFound           = rpccode.Code_NOT_FOUND
	AlreadyExists      = rpccode.Code_ALREADY_EXISTS
	PermissionDenied   = rpccode.Code_PERMISSION_DENIED
	ResourceExhausted  = rpccode.Code_RESOURCE_EXHAUSTED
	FailedPrecondition = rpccode.Code_FAILED_PRECONDITION
	Aborted            = rpccode.Code_ABORTED
	OutOfRange         = rpccode.Code_OUT_OF_RANGE
	Unimplemented      = rpccode.Code_UNIMPLEMENTED
	Internal           = rpccode.Code_INTERNAL
	Unavailable        = rpccode.Code_UNAVAILABLE
	DataLoss           = rpccode.Code_DATA_LOSS
	Unauthenticated    = rpccode.Code_UNAUTHENTICATED
)

func New(msg string) error {
	return internal.NewCanonicalError(Unknown, msg, nil)
}

func Error(c rpccode.Code, msg string) error {
	return internal.NewCanonicalError(c, msg, nil)
}

func Errorf(c rpccode.Code, format string, a ...interface{}) error {
	fmtErr := fmt.Errorf(format, a...)
	wrapped := errors.Unwrap(fmtErr)
	return internal.NewCanonicalError(c, fmtErr.Error(), wrapped)
}

func Code(err error) rpccode.Code {
	if err == nil {
		return Ok
	}

	if ce, ok := err.(*internal.CanonicalError); ok {
		return ce.Code()
	}
	return Unknown
}

func CodeAsString(err error) string {
	return Code(err).String()
}

func IsOK(err error) bool                 { return err == nil }
func IsCancelled(err error) bool          { return Code(err) == Cancelled }
func IsUnknown(err error) bool            { return Code(err) == Unknown }
func IsInvalidArgument(err error) bool    { return Code(err) == InvalidArgument }
func IsDeadlineExceeded(err error) bool   { return Code(err) == DeadlineExceed }
func IsNotFound(err error) bool           { return Code(err) == NotFound }
func IsAlreadyExists(err error) bool      { return Code(err) == AlreadyExists }
func IsPermissionDenied(err error) bool   { return Code(err) == PermissionDenied }
func IsResourceExhausted(err error) bool  { return Code(err) == ResourceExhausted }
func IsFailedPrecondition(err error) bool { return Code(err) == FailedPrecondition }
func IsAborted(err error) bool            { return Code(err) == Aborted }
func IsOutOfRange(err error) bool         { return Code(err) == OutOfRange }
func IsUnimplemented(err error) bool      { return Code(err) == Unimplemented }
func IsInternal(err error) bool           { return Code(err) == Internal }
func IsUnavailable(err error) bool        { return Code(err) == Unavailable }
func IsDataLoss(err error) bool           { return Code(err) == DataLoss }
func IsUnauthenticated(err error) bool    { return Code(err) == Unauthenticated }
