package websocket

import (
	"context"
	"errors"
	"fmt"
	"github.com/TemaKut/messenger-apigateway/internal/app/adapter/auth"
	"github.com/TemaKut/messenger-apigateway/internal/app/logger"
	delegatedto "github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
	"github.com/google/uuid"
	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"sync"
)

type Session struct {
	id string

	conn *websocket.Conn

	delegateService DelegateService

	states       map[sessionStateType]sessionState
	currentState sessionState

	doneCh chan struct{}
	wg     sync.WaitGroup

	logger *logger.Logger
}

func NewSession(
	conn *websocket.Conn,
	delegateService DelegateService,
	logger *logger.Logger,
) *Session {
	session := &Session{
		id:              uuid.NewString(),
		conn:            conn,
		delegateService: delegateService,
		doneCh:          make(chan struct{}, 1),
		logger:          logger,
	}

	unauthorizedState := newUnauthorizedSessionState(session)

	session.states = map[sessionStateType]sessionState{
		sessionStateTypeUnauthorized: unauthorizedState,
		sessionStateTypeAuthorized:   newAuthorizedSessionState(session),
	}

	session.currentState = unauthorizedState

	return session
}

func (s *Session) HandleRequests(ctx context.Context) error {
mainCycle:
	for {
		select {
		case <-s.doneCh:
			s.logger.Debugf("session handle request cancelled")

			break mainCycle
		default:
		}

		reqBytes := make([]byte, 0)

		if err := websocket.Message.Receive(s.conn, &reqBytes); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return fmt.Errorf("error receive message. %w", err)
		}

		var req pb.Request

		if err := proto.Unmarshal(reqBytes, &req); err != nil {
			return fmt.Errorf("error unmarshalling request. %w", err)
		}

		s.wg.Add(1)

		err := func() error {
			defer s.wg.Done()

			if req.GetId() == "" {
				err := s.sendResponse(req.GetId(), nil, s.encodeResponseErrorSource(errRequestHasNoId))
				if err != nil {
					return fmt.Errorf("error sending response. %w", err)
				}
			}

			if err := s.currentState.handleRequest(s.ctx(), &req); err != nil {
				return fmt.Errorf("error handling request. %w", err)
			}

			return nil
		}()
		if err != nil {
			return fmt.Errorf("error process request. %w", err)
		}
	}

	return nil
}

func (s *Session) Shutdown() {
	s.logger.Debugf("session shutdown")

	close(s.doneCh)
	s.wg.Wait()
}

func (s *Session) handleRequest(ctx context.Context, req *pb.Request) error {
	var (
		success       *pb.Success
		responseError *pb.Error
	)

	switch { // TODO мне не нравится подобная обработка внутри кейса. Этот метод сильно разрастётся
	case req.GetUserRegister() != nil:
		resp, err := s.delegateService.OnUserRegisterRequest(ctx, decodeUserRegisterRequest(req.GetUserRegister()))
		if err != nil {
			responseError = s.encodeResponseErrorSource(err)
			break
		}

		success = encodeUserRegisterResponse(resp)
	case req.GetUserAuthorize() != nil:
		requestDecoded, err := decodeUserAuthorizeRequest(req.GetUserAuthorize())
		if err != nil {
			responseError = s.encodeResponseErrorSource(err)
			break
		}

		resp, err := s.delegateService.OnUserAuthorizeRequest(ctx, requestDecoded)
		if err != nil {
			responseError = s.encodeResponseErrorSource(err)
			break
		}

		success = encodeUserAuthorizeResponse(resp)
	default:
		return fmt.Errorf("error unsupported request type")
	}

	if err := s.sendResponse(req.GetId(), success, responseError); err != nil {
		return fmt.Errorf("error sending response. %w", err)
	}

	return nil
}

func (s *Session) setState(stateType sessionStateType) error {
	state, ok := s.states[stateType]
	if !ok {
		return fmt.Errorf("session state type %d not found", stateType)
	}

	s.currentState = state

	return nil
}

func (s *Session) ctx() context.Context { // TODO данные о юзере и тд. обработку контекста скорее всего в либу
	ctx := context.Background()

	ctx = context.WithValue(ctx, "Key", "Value")

	return ctx
}

func (s *Session) sendResponse(id string, successSource *pb.Success, errorSource *pb.Error) error {
	response := &pb.Response{
		Id:         id,
		ServerTime: timestamppb.Now(),
	}

	switch {
	case successSource != nil:
		response.Source = &pb.Response_Success{
			Success: successSource,
		}
	case errorSource != nil:
		response.Source = &pb.Response_Error{
			Error: errorSource,
		}
	default:
		return fmt.Errorf("error response (id=%s) has no source", id)
	}

	protoBytes, err := proto.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshalling response. %w", err)
	}

	if err := websocket.Message.Send(s.conn, protoBytes); err != nil {
		return fmt.Errorf("error sending response. %w", err)
	}

	return nil
}

func (s *Session) encodeResponseErrorSource(err error) *pb.Error {
	s.logger.Errorf("error for session (id=%s). %s", s.id, err)

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
	case errors.Is(err, auth.ErrInvalidCredentials):
		errorMessage = pb.Error{
			Reason:      pb.ErrorReason_ERROR_REASON_USER_INVALID_CREDENTIALS,
			Description: "invalid user credentials",
		}
	default:
		errorMessage = pb.Error{
			Reason:      pb.ErrorReason_ERROR_REASON_UNKNOWN,
			Description: "unknown error",
		}
	}

	return &errorMessage
}
