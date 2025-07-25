package auth

import "errors"

var (
	ErrUnknown                = errors.New("unknown")
	ErrUserEmailAlreadyExists = errors.New("error-user-email-already-exists")
)
