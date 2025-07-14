package websocket

import (
	"github.com/TemaKut/messenger-apigateway/internal/app/handler/websocket/controllers"
	"github.com/TemaKut/messenger-apigateway/internal/app/logger"
	"golang.org/x/net/websocket"
)

type Handler struct {
	logger *logger.Logger

	controller *controllers.Controller
}

func NewHandler(
	controller *controllers.Controller,
	logger *logger.Logger,
) *Handler {
	return &Handler{
		controller: controller,
		logger:     logger,
	}
}

func (h *Handler) Handle(conn *websocket.Conn) {
	h.logger.Debugf("new websocket connection")

	defer func() {
		h.logger.Debugf("close websocket connection")

		_ = conn.Close()
	}()

	session := NewSession(conn, h.controller, h.logger)
	defer session.Shutdown()

	if err := session.HandleRequests(conn.Request().Context()); err != nil {
		h.logger.Errorf("error handling requests: %s", err)
	}
}
