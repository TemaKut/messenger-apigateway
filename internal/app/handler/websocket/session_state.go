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
	case req.GetUserAuthorize() != nil:
	default:
		err := fmt.Errorf("error unsupported unauthorized request (id=%s). %w", req.GetId(), errForbidden)
		if err := u.session.sendResponse(req.GetId(), nil, u.session.encodeResponseErrorSource(err)); err != nil {
			return fmt.Errorf("error sending response. %w", err)
		}

		return nil
	}

	if err := u.session.handleRequest(ctx, req); err != nil {
		return fmt.Errorf("error handle request. %w", err)
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
	//case req.GetUserAuthorize() != nil: // TODO конкретная ошибка о том что уже была авторизация
	default:
		return fmt.Errorf("error unsupported request type")
	}
}
