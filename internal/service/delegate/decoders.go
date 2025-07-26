package delegatesvc

import (
	"fmt"
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
	delegatedto "github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
)

func decodeUserRegisterRequest(req delegatedto.UserRegisterRequest) authdto.RegisterUserRequest {
	return authdto.RegisterUserRequest{
		Name:     req.Name,
		LastName: req.LastName,
		Email:    req.Email,
		Password: req.Password,
	}
}

func decodeUserAuthorizeRequest(req delegatedto.UserAuthorizeRequest) (authdto.UserAuthorizeRequest, error) {
	credentialsDecoded, err := decodeUserAuthorizeCredentials(req.Credentials)
	if err != nil {
		return authdto.UserAuthorizeRequest{}, fmt.Errorf("error decode user authorize credentials. %w", err)
	}

	return authdto.UserAuthorizeRequest{
		Credentials: credentialsDecoded,
	}, nil
}

func decodeUserAuthorizeCredentials(
	credentials delegatedto.UserAuthorizeCredentials,
) (authdto.UserAuthorizeCredentials, error) {
	var credentialsDecoded authdto.UserAuthorizeCredentials

	switch {
	case credentials.Email != nil:
		credentialsDecoded.Email = toPtr(decodeUserAuthorizeEmailCredentials(*credentials.Email))
	default:
		return authdto.UserAuthorizeCredentials{}, fmt.Errorf("credentials not passed")
	}

	return credentialsDecoded, nil
}

func decodeUserAuthorizeEmailCredentials(
	credential delegatedto.UserAuthorizeEmailCredential,
) authdto.UserAuthorizeEmailCredential {
	return authdto.UserAuthorizeEmailCredential{
		Email:    credential.Email,
		Password: credential.Password,
	}
}

func toPtr[T any](ptr T) *T {
	return &ptr
}
