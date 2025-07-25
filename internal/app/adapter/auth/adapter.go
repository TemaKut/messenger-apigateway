package auth

import (
	"context"
	"fmt"
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
	authv1 "github.com/TemaKut/messenger-service-proto/gen/go/auth"
)

type Adapter struct {
	userApiClient authv1.UserAPIClient
}

func NewAdapter(userApiClient authv1.UserAPIClient) *Adapter {
	return &Adapter{
		userApiClient: userApiClient,
	}
}

func (a *Adapter) RegisterUser(
	ctx context.Context,
	req authdto.RegisterUserRequest,
) (authdto.RegisterUserResponse, error) {
	userRegisterResponse, err := a.userApiClient.UserRegister(ctx, &authv1.UserAPIUserRegisterRequest{
		Name:     req.Name,
		LastName: req.LastName,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return authdto.RegisterUserResponse{}, fmt.Errorf("error registering user. %w", encodeError(err))
	}

	return authdto.RegisterUserResponse{
		User: encodeUser(userRegisterResponse.GetUser()),
	}, nil
}
