package websocket

import (
	"context"
	"errors"
	"fmt"
	"github.com/TemaKut/messenger-apigateway/internal/app/logger"
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

		requestBytes := make([]byte, 0)

		if err := websocket.Message.Receive(s.conn, &requestBytes); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return fmt.Errorf("error receive message. %w", err)
		}

		var req pb.Request

		if err := proto.Unmarshal(requestBytes, &req); err != nil {
			return fmt.Errorf("error unmarshalling request. %w", err)
		}

		if req.GetId() == "" {
			if err := s.sendError("", errRequestHasNoId); err != nil {
				return fmt.Errorf("error send error. %w", err)
			}

			continue
		}

		s.wg.Add(1)

		err := func() error {
			defer s.wg.Done()

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
	switch {
	case req.GetUserRegister() != nil:
		resp, err := s.delegateService.OnUserRegisterRequest(ctx, decodeUserRegisterRequest(req.GetUserRegister()))
		if err != nil {
			return s.sendError(req.GetId(), err)
		}

		if err := s.sendResponse(req.GetId(), encodeUserRegisterResponse(resp)); err != nil {
			return fmt.Errorf("error sending response. %w", err)
		}
	case req.GetUserAuthorize() != nil:
		requestDecoded, err := decodeUserAuthorizeRequest(req.GetUserAuthorize())
		if err != nil {
			return s.sendError(req.GetId(), err)
		}

		resp, err := s.delegateService.OnUserAuthorizeRequest(ctx, requestDecoded)
		if err != nil {
			return s.sendError(req.GetId(), err)
		}

		if err := s.sendResponse(req.GetId(), encodeUserAuthorizeResponse(resp)); err != nil {
			return fmt.Errorf("error sending response. %w", err)
		}
	default:
		return fmt.Errorf("error unsupported request type")
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

func (s *Session) sendResponse(requestId string, responseData proto.Message) error {
	var response pb.Response

	switch data := responseData.(type) {
	case *pb.UserRegisterResponse:
		response.Data = &pb.Response_UserRegister{UserRegister: data}
	case *pb.UserAuthorizeResponse:
		response.Data = &pb.Response_UserAuthorize{UserAuthorize: data}
	default:
		return fmt.Errorf("error unknown response type: %T", data)
	}

	responseContainer := pb.ServerMessageContainer{
		RequestId:  requestId,
		ServerTime: timestamppb.Now(),
		Data: &pb.ServerMessageContainer_Response{
			Response: &response,
		},
	}

	protoBytes, err := proto.Marshal(&responseContainer)
	if err != nil {
		return fmt.Errorf("error marshalling server message container. %w", err)
	}

	if err := websocket.Message.Send(s.conn, protoBytes); err != nil {
		return fmt.Errorf("error sending response. %w", err)
	}

	return nil
}

func (s *Session) sendError(requestId string, errForSend error) error {
	s.logger.Errorf("error for session (id=%s). %s", s.id, errForSend)

	errorContainer := pb.ServerMessageContainer{
		RequestId:  requestId,
		ServerTime: timestamppb.Now(),
		Data: &pb.ServerMessageContainer_Error{
			Error: encodeError(errForSend),
		},
	}

	protoBytes, err := proto.Marshal(&errorContainer)
	if err != nil {
		return fmt.Errorf("error marshalling server message container. %w", err)
	}

	if err := websocket.Message.Send(s.conn, protoBytes); err != nil {
		return fmt.Errorf("error sending error. %w", err)
	}

	return nil
}
