package otp

import "errors"

var (
	ErrNil      error = errors.New("otp: Has not been initialized")
	ErrID             = errors.New("otp: Error id")
	ErrType           = errors.New("otp: Error type")
	ErrExist          = errors.New("otp: Exist id")
	ErrNotExist       = errors.New("otp: Not exist id")
	ErrNotOpen        = errors.New("otp: Close this auth")
)
