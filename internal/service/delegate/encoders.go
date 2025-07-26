package delegatesvc

import (
	"errors"
	"fmt"
	"github.com/TemaKut/messenger-apigateway/internal/app/adapter/auth"
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
	delegatedto "github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
)

func encodeUserRegisterResponse(response authdto.RegisterUserResponse) delegatedto.UserRegisterResponse {
	return delegatedto.UserRegisterResponse{
		User: encodeUser(response.User),
	}
}

func encodeUser(user authdto.User) delegatedto.User {
	return delegatedto.User{
		Id:       user.Id,
		Name:     user.Name,
		LastName: user.LastName,
	}
}

func encodeAuthorizeResponse(response authdto.UserAuthorizeResponse) delegatedto.UserAuthorizeResponse {
	return delegatedto.UserAuthorizeResponse{
		User:       encodeUser(response.User),
		AuthParams: encodeAuthParams(response.AuthParams),
	}
}

func encodeAuthParams(params authdto.AuthParams) delegatedto.AuthParams {
	return delegatedto.AuthParams{
		AccessToken:  encodeAuthToken(params.AccessToken),
		RefreshToken: encodeAuthToken(params.RefreshToken),
	}
}

func encodeAuthToken(token authdto.AuthToken) delegatedto.AuthToken {
	return delegatedto.AuthToken{
		Token:     token.Token,
		ExpiredAt: token.ExpiredAt,
	}
}

func encodeError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, auth.ErrUserEmailAlreadyExists):
		return fmt.Errorf("%w, %w", delegatedto.ErrUserEmailAlreadyExists, err)
	case errors.Is(err, auth.ErrInvalidCredentials):
		return fmt.Errorf("%w. %w", delegatedto.ErrInvalidUserCredentials, err)
	default:
		return fmt.Errorf("%w, %w", delegatedto.ErrUnknown, err)
	}
}
