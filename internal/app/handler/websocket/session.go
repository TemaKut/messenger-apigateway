package websocket

import (
	"context"
	"fmt"
	"github.com/TemaKut/messenger-apigateway/internal/app/logger"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
	"io"
	"sync"
)

type Session struct {
	conn *websocket.Conn

	states       map[sessionStateType]sessionState
	currentState sessionState

	doneCh chan struct{}
	wg     sync.WaitGroup

	logger *logger.Logger

	// Depends  TODO не лайкаю чтобы именно сессия делегировала запросы. Изучить и применить другое решение
	authService AuthService
	// ^^^^^^^^^^
}

func NewSession(
	conn *websocket.Conn,
	authService AuthService,
	logger *logger.Logger,
) *Session {
	session := &Session{
		conn:        conn,
		authService: authService,
		doneCh:      make(chan struct{}, 1),
		logger:      logger,
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

		s.wg.Add(1)

		err := func() error {
			defer s.wg.Done()

			reqBytes := make([]byte, 0)

			if err := websocket.Message.Receive(s.conn, &reqBytes); err != nil {
				if err == io.EOF {
					return nil
				}

				return fmt.Errorf("error receive message. %w", err)
			}

			var req pb.Request

			if err := proto.Unmarshal(reqBytes, &req); err != nil {
				return fmt.Errorf("error unmarshalling request. %w", err)
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

func (s *Session) sendResponse(ctx context.Context, response *pb.Response) error {
	protoBytes, err := proto.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshalling response. %w", err)
	}

	if err := websocket.Message.Send(s.conn, protoBytes); err != nil {
		return fmt.Errorf("error sending response. %w", err)
	}

	return nil
}

func (s *Session) getAuthService() AuthService {
	return s.authService
}
