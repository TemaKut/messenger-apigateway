package factory

import (
	"github.com/TemaKut/messenger-apigateway/internal/app/handler/websocket"
	delegatesvc "github.com/TemaKut/messenger-apigateway/internal/service/delegate"
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	delegatesvc.NewService,
	wire.Bind(new(websocket.DelegateService), new(*delegatesvc.Service)),
)
