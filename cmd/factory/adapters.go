package factory

import (
	"github.com/TemaKut/messenger-apigateway/internal/app/adapter/auth"
	delegatesvc "github.com/TemaKut/messenger-apigateway/internal/service/delegate"
	"github.com/google/wire"
)

var AdaptersSet = wire.NewSet(
	auth.NewAdapter,
	wire.Bind(new(delegatesvc.AuthService), new(*auth.Adapter)),
)
