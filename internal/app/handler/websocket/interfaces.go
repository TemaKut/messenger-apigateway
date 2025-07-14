package websocket

import (
	"context"
	"github.com/TemaKut/messenger-apigateway/internal/dto/auth"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
)

type sessionState interface {
	handleRequest(ctx context.Context, req *pb.Request) error
}
type AuthService interface {
	RegisterUser(
		ctx context.Context,
		req authdto.RegisterUserRequest,
	) (authdto.RegisterUserResponse, error)
}
