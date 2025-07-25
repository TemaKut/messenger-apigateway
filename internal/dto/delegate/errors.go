package delegatedto

import "errors"

var (
	ErrUnknown                = errors.New("unknown")
	ErrUserEmailAlreadyExists = errors.New("user-email-already-exists")
)
