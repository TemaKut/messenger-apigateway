package websocket

import (
	"context"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
)

type sessionState interface {
	handleRequest(ctx context.Context, req *pb.Request) error
}
