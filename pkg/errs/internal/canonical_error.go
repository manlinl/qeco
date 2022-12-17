package internal

import (
	"log"

	rpccode "google.golang.org/genproto/googleapis/rpc/code"
)

type CanonicalError struct {
	code      rpccode.Code
	msg       string
	origError error
}

func NewCanonicalError(c rpccode.Code, msg string, origErr error) *CanonicalError {
	if c == rpccode.Code_OK {
		log.Fatal("expect non-OK error code")
	}

	return &CanonicalError{
		code:      c,
		msg:       msg,
		origError: origErr,
	}
}

func (e *CanonicalError) Error() string {
	return e.msg
}

func (e *CanonicalError) Unwrap() error {
	return e.origError
}

func (e *CanonicalError) Code() rpccode.Code {
	return e.code
}
