package websocket

type sessionStateType int

const (
	sessionStateTypeUnauthorized sessionStateType = iota
	sessionStateTypeAuthorized
)
