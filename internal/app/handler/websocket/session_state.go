package websocket

import (
	"context"
	"fmt"
)

type unauthorizedSessionState struct {
	session *Session
}

func newUnauthorizedSessionState(session *Session) *unauthorizedSessionState {
	return &unauthorizedSessionState{session: session}
}

func (u *unauthorizedSessionState) handleRequest(ctx context.Context, req *Request) error {
	switch {
	case req.Request().GetAuth() != nil:
		return u.handleAuthRequest(ctx, req.GetAuth())
	default:
		return fmt.Errorf("error unsupported request type")
	}

	return nil
}

type authorizedSessionState struct {
	session *Session
}

func newAuthorizedSessionState(session *Session) *authorizedSessionState {
	return &authorizedSessionState{session: session}
}

func (u *authorizedSessionState) handleRequest(ctx context.Context, req *Request) error {
	switch {
	case req.Request().GetAuth() != nil:
		return u.session.invoke(u.session.getController().Auth.UserRegisterController.Invoke, req)
	default:
		return fmt.Errorf("error unsupported request type")
	}
}
