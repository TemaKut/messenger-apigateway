package delegatedto

import "errors"

var (
	ErrUnknown                = errors.New("unknown")
	ErrUserEmailAlreadyExists = errors.New("user-email-already-exists")
	ErrInvalidUserCredentials = errors.New("invalid-user-credentials")
	ErrValidation             = errors.New("validation")
)
