package auth

import (
	"fmt"
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
	authv1 "github.com/TemaKut/messenger-service-proto/gen/go/auth"
)

func decodeUserAuthorizeRequest(req authdto.UserAuthorizeRequest) (*authv1.UserAPIAuthorizeRequest, error) {
	var authorizeRequest authv1.UserAPIAuthorizeRequest

	switch {
	case req.Credentials.Email != nil:
		authorizeRequest.Credentials = &authv1.UserAPIAuthorizeRequest_Email{
			Email: decodeUserAuthorizeEmailCredentials(*req.Credentials.Email),
		}
	default:
		return nil, fmt.Errorf("error credentials not passed")
	}

	return &authorizeRequest, nil
}

func decodeUserAuthorizeEmailCredentials(
	credential authdto.UserAuthorizeEmailCredential,
) *authv1.UserAPIAuthorizeEmailCredentials {
	return &authv1.UserAPIAuthorizeEmailCredentials{
		Email:    credential.Email,
		Password: credential.Password,
	}
}
