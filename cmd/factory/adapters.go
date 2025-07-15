package factory

import (
	"github.com/TemaKut/messenger-apigateway/internal/app/adapter/auth"
	"github.com/TemaKut/messenger-apigateway/internal/app/handler/websocket"
	"github.com/google/wire"
)

var AdaptersSet = wire.NewSet(
	auth.NewAdapter,
	wire.Bind(new(websocket.AuthService), new(*auth.Adapter)),
)
