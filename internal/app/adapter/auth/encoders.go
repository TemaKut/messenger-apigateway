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
	case "user-email-already-exist":
		return fmt.Errorf("%w. %s", ErrUserEmailAlreadyExists, st.Message())
	default:
		return ErrUnknown
	}
}
