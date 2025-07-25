package websocket

import "errors"

var (
	errRequestHasNoId = errors.New("request has no id")
	errForbidden      = errors.New("forbidden")
)
