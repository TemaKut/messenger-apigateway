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
	var success *pb.Success

	switch { // TODO этот метод будет сильно разростаться подумать как его компактизировать
	case req.GetUserRegister() != nil:
		resp, err := u.session.getDelegateService().OnUserRegisterRequest(
			ctx,
			decodeUserRegisterRequest(req.GetUserRegister()),
		)
		if err != nil {
			return fmt.Errorf("error delegate user register request. %w", err)
		}

		success = encodeUserRegisterResponse(resp)
	default:
		return fmt.Errorf("error unsupported request type")
	}

	if err := u.session.sendSuccessResponse(req.GetId(), success); err != nil {
		return fmt.Errorf("error sending success response. %w", err)
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
