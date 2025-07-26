package auth

import "errors"

var (
	ErrUnknown                = errors.New("unknown")
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
	ErrInvalidCredentials     = errors.New("invalid user credentials")
)
