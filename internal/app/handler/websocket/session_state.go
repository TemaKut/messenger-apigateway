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
	case req.GetAuth() != nil:
		//return u.handleAuthRequest(ctx, req.GetAuth())
	default:
		return fmt.Errorf("error unsupported request type")
	}
}

type authorizedSessionState struct {
	session *Session
}

func newAuthorizedSessionState(session *Session) *authorizedSessionState {
	return &authorizedSessionState{session: session}
}

func (u *authorizedSessionState) handleRequest(ctx context.Context, req *pb.Request) error {
	switch {
	case req.GetAuth() != nil:
		return nil // TODO тут другой набор реквестов
	default:
		return fmt.Errorf("error unsupported request type")
	}
}
