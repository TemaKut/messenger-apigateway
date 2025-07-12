package factory

import (
	"github.com/TemaKut/messenger-apigateway/internal/app/logger"
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	ProvideApp,
	ProvideLogger,
)

type App struct{}

func ProvideApp(
	logger *logger.Logger,
	_ HttpProvider,
) (App, func()) {
	logger.Infof("app inited")

	return App{}, func() {
		logger.Infof("app shutting down")
	}
}

func ProvideLogger() (*logger.Logger, error) {
	return logger.NewLogger(logger.LogLevelDebug) // TODO из конфига определить
}
