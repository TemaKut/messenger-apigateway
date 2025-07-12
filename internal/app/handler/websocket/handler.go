package websocket

import (
	"fmt"
	"github.com/TemaKut/messenger-apigateway/internal/app/logger"
	"golang.org/x/net/websocket"
	"io"
)

type Handler struct {
	logger *logger.Logger
}

func NewHandler(logger *logger.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) Handle(conn *websocket.Conn) {
	h.logger.Debugf("new websocket connection")

	defer func() {
		h.logger.Debugf("close websocket connection")

		if err := conn.Close(); err != nil {
			h.logger.Errorf("error closing websocket connection: %s", err)
		}
	}()

	for {
		var message string

		if err := websocket.Message.Receive(conn, &message); err != nil {
			if err != io.EOF {
				h.logger.Errorf("error receiving message: %s", err)
			}

			break
		}

		fmt.Println(message)
	}
}
