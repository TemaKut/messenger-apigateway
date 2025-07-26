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

func (a *Adapter) Register(
	ctx context.Context,
	req authdto.RegisterUserRequest,
) (authdto.RegisterUserResponse, error) {
	userRegisterResponse, err := a.userApiClient.Register(ctx, &authv1.UserAPIRegisterRequest{
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

func (a *Adapter) Authorize(
	ctx context.Context,
	req authdto.UserAuthorizeRequest,
) (authdto.UserAuthorizeResponse, error) {
	authorizeRequest, err := decodeUserAuthorizeRequest(req)
	if err != nil {
		return authdto.UserAuthorizeResponse{}, fmt.Errorf("error decodeauthorize request. %w", err)
	}

	authorizeResponse, err := a.userApiClient.Authorize(ctx, authorizeRequest)
	if err != nil {
		return authdto.UserAuthorizeResponse{}, fmt.Errorf("error authorize. %w", encodeError(err))
	}

	return encodeAuthorizeResponse(authorizeResponse), nil
}
