package auth

import (
	"fmt"
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
	authv1 "github.com/TemaKut/messenger-service-proto/gen/go/auth"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func encodeUser(user *authv1.User) authdto.User {
	return authdto.User{
		Id:       user.GetId(),
		Name:     user.GetName(),
		LastName: user.GetLastName(),
	}
}

func encodeAuthorizeResponse(response *authv1.UserAPIAuthorizeResponse) authdto.UserAuthorizeResponse {
	return authdto.UserAuthorizeResponse{
		User:       encodeUser(response.GetUser()),
		AuthParams: encodeAuthParams(response.GetAuthParams()),
	}
}

func encodeAuthParams(params *authv1.AuthParams) authdto.AuthParams {
	return authdto.AuthParams{
		AccessToken:  encodeAuthToken(params.GetAccessToken()),
		RefreshToken: encodeAuthToken(params.GetRefreshToken()),
	}
}

func encodeAuthToken(token *authv1.AuthToken) authdto.AuthToken {
	return authdto.AuthToken{
		Token:     token.GetToken(),
		ExpiredAt: token.GetExpiredAt().AsTime(),
	}
}

func encodeError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return nil
	}

	var detailReason string

	for _, detail := range st.Details() {
		switch detailType := detail.(type) {
		case *errdetails.ErrorInfo:
			detailReason = detailType.GetReason()
		}
	}

	switch detailReason {
	case "auth.user-email-already-exist":
		return fmt.Errorf("%w. %s", ErrUserEmailAlreadyExists, st.Message())
	case "auth.invalid-user-credentials":
		return fmt.Errorf("%w. %s", ErrInvalidCredentials, st.Message())
	default:
		return ErrUnknown
	}
}
