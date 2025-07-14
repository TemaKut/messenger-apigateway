package websocket

import (
	"context"
	"fmt"
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
		return u.handleUserRegisterRequest(ctx, req.GetUserRegister())
	default:
		return fmt.Errorf("error unsupported request type")
	}
}

func (u *unauthorizedSessionState) handleUserRegisterRequest(
	ctx context.Context,
	register *pb.UserRegisterRequest,
) error {
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
