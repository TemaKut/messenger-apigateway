package factory

import (
	"github.com/TemaKut/messenger-apigateway/internal/app/handler/websocket/controllers"
	"github.com/google/wire"
)

var ControllersSet = wire.NewSet(
	controllers.NewController,
	
	controllers.NewUserRegisterControllerImpl,
	wire.Bind(new(controllers.UserRegisterController), new(*controllers.UserRegisterControllerImpl)),
)
