package websocket

import (
	"context"
	"fmt"
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
)

type unauthorizedSessionState struct {
	session *Session
}

func newUnauthorizedSessionState(session *Session) *unauthorizedSessionState {
	return &unauthorizedSessionState{session: session}
}

func (u *unauthorizedSessionState) handleRequest(ctx context.Context, req *pb.Request) error {
	switch {
	case req.GetUserRegister() != nil:
		return u.handleUserRegisterRequest(ctx, req)
	default:
		return fmt.Errorf("error unsupported request type")
	}
}

func (u *unauthorizedSessionState) handleUserRegisterRequest(
	ctx context.Context,
	req *pb.Request,
) error {
	userRegisterRequest := req.GetUserRegister()

	registerUserResponse, err := u.session.getAuthService().RegisterUser(ctx, authdto.RegisterUserRequest{
		Name:     userRegisterRequest.GetName(),
		LastName: userRegisterRequest.GetLastName(),
		Email:    userRegisterRequest.GetEmail(),
		Password: userRegisterRequest.GetPassword(),
	})
	if err != nil {
		return fmt.Errorf("error register user. %w", err)
	}

	err = u.session.sendResponse(ctx, &pb.Response{
		Id: req.GetId(),
		Source: &pb.Response_Success{
			Success: &pb.Success{
				Data: &pb.Success_UserRegister{
					UserRegister: &pb.UserRegisterResponse{
						User: encodeUser(registerUserResponse.User),
					},
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("error send response. %w", err)
	}

	return nil
}

type authorizedSessionState struct {
	session *Session
}

func newAuthorizedSessionState(session *Session) *authorizedSessionState {
	return &authorizedSessionState{session: session}
}

func (u *authorizedSessionState) handleRequest(ctx context.Context, req *pb.Request) error {
	switch { // TODO тут запросы которые требуют авторизации
	//case req.Request().GetAuth() != nil:
	//	return u.session.invoke(u.session.getController().Auth.UserRegisterController.Invoke, req)
	default:
		return fmt.Errorf("error unsupported request type")
	}
}
