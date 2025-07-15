package delegatesvc

import (
	"context"
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
)

type AuthService interface {
	RegisterUser(
		ctx context.Context,
		req authdto.RegisterUserRequest,
	) (authdto.RegisterUserResponse, error)
}
