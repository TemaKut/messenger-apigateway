package websocket

import (
	"errors"
	delegatedto "github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
	pbuser "github.com/TemaKut/messenger-client-proto/gen/go/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func encodeUserRegisterResponse(resp delegatedto.UserRegisterResponse) *pb.UserRegisterResponse {
	return &pb.UserRegisterResponse{
		User: encodeUser(resp.User),
	}
}

func encodeUserAuthorizeResponse(resp delegatedto.UserAuthorizeResponse) *pb.UserAuthorizeResponse {
	return &pb.UserAuthorizeResponse{
		User:       encodeUser(resp.User),
		AuthParams: encodeAuthParams(resp.AuthParams),
	}
}

func encodeAuthParams(params delegatedto.AuthParams) *pbuser.AuthParams {
	return &pbuser.AuthParams{
		AccessToken:  encodeAuthToken(params.AccessToken),
		RefreshToken: encodeAuthToken(params.RefreshToken),
	}
}

func encodeAuthToken(token delegatedto.AuthToken) *pbuser.AuthToken {
	return &pbuser.AuthToken{
		Token:     token.Token,
		ExpiredAt: timestamppb.New(token.ExpiredAt),
	}
}

func encodeUser(u delegatedto.User) *pbuser.User {
	return &pbuser.User{
		Id:       u.Id,
		Name:     u.Name,
		LastName: u.LastName,
	}
}

func encodeError(err error) *pb.Error {
	var errorMessage pb.Error

	switch {
	case errors.Is(err, delegatedto.ErrUserEmailAlreadyExists):
		errorMessage = pb.Error{
			Reason:      pb.ErrorReason_ERROR_REASON_USER_EMAIL_ALREADY_EXISTS,
			Description: "user email already exists",
		}
	case errors.Is(err, errRequestHasNoId):
		errorMessage = pb.Error{
			Reason:      pb.ErrorReason_ERROR_REASON_REQUEST_HAS_NO_ID,
			Description: "request has no identifier",
		}
	case errors.Is(err, errForbidden):
		errorMessage = pb.Error{
			Reason:      pb.ErrorReason_ERROR_REASON_FORBIDDEN,
			Description: "forbidden",
		}
	case errors.Is(err, delegatedto.ErrInvalidUserCredentials):
		errorMessage = pb.Error{
			Reason:      pb.ErrorReason_ERROR_REASON_USER_INVALID_CREDENTIALS,
			Description: "invalid user credentials",
		}
	case errors.Is(err, delegatedto.ErrValidation):
		errorMessage = pb.Error{
			Reason:      pb.ErrorReason_ERROR_REASON_VALIDATION,
			Description: "validation error",
		}
	default:
		errorMessage = pb.Error{
			Reason:      pb.ErrorReason_ERROR_REASON_UNKNOWN,
			Description: "unknown error",
		}
	}

	return &errorMessage
}
