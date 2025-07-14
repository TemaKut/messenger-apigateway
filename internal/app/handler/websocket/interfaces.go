package websocket

import (
	"context"
)

type sessionState interface {
	handleRequest(ctx context.Context, req *Request) error
}
