package websocket

import (
	"fmt"
	delegatedto "github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
)

func decodeUserRegisterRequest(req *pb.UserRegisterRequest) delegatedto.UserRegisterRequest {
	return delegatedto.UserRegisterRequest{
		Name:     req.GetName(),
		LastName: req.GetLastName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func decodeUserAuthorizeRequest(req *pb.UserAuthorizeRequest) (delegatedto.UserAuthorizeRequest, error) {
	delegateRequest := delegatedto.UserAuthorizeRequest{}

	switch {
	case req.GetEmail() != nil:
		delegateRequest.Credentials.Email = toPtr(decodeUserAuthorizeEmailCredentials(req.GetEmail()))
	default:
		return delegatedto.UserAuthorizeRequest{}, fmt.Errorf("error request has no credentials")
	}

	return delegateRequest, nil
}

func decodeUserAuthorizeEmailCredentials(
	credentials *pb.UserAuthorizeEmailCredential,
) delegatedto.UserAuthorizeEmailCredential {
	return delegatedto.UserAuthorizeEmailCredential{
		Email:    credentials.GetEmail(),
		Password: credentials.GetPassword(),
	}
}

func toPtr[T any](v T) *T {
	return &v
}
