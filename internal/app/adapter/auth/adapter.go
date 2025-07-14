package auth

import (
	"context"
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) RegisterUser(
	ctx context.Context,
	req authdto.RegisterUserRequest,
) (authdto.RegisterUserResponse, error) {
	return authdto.RegisterUserResponse{}, nil
}
