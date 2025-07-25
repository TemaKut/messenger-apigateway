package websocket

import (
	"context"
	"github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
)

type sessionState interface {
	handleRequest(ctx context.Context, req *pb.Request) error
}
type DelegateService interface {
	OnUserRegisterRequest(
		ctx context.Context,
		req delegatedto.UserRegisterRequest,
	) (delegatedto.UserRegisterResponse, error)
	OnUserAuthorizeRequest(
		ctx context.Context,
		req delegatedto.UserAuthorizeRequest,
	) (delegatedto.UserAuthorizeResponse, error)
}
