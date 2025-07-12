package websocket

import (
	"github.com/TemaKut/messenger-apigateway/internal/app/logger"
	"golang.org/x/net/websocket"
)

type Handler struct {
	logger *logger.Logger
}

func NewHandler(logger *logger.Logger) *Handler {
	return &Handler{}
}

func (h *Handler) Handle(conn *websocket.Conn) {
	h.logger.Debugf("new websocket connection")

	defer func() {
		h.logger.Debugf("close websocket connection")

		if err := conn.Close(); err != nil {
			h.logger.Errorf("error closing websocket connection: %s", err)
		}
	}()
}
