package delegatesvc

import (
	"context"
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
)

type AuthService interface {
	Register(ctx context.Context, req authdto.RegisterUserRequest) (authdto.RegisterUserResponse, error)
	Authorize(ctx context.Context, req authdto.UserAuthorizeRequest) (authdto.UserAuthorizeResponse, error)
}
