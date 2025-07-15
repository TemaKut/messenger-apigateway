package factory

import (
	"fmt"
	"github.com/TemaKut/messenger-apigateway/internal/app/config"
	"github.com/TemaKut/messenger-apigateway/internal/app/handler/websocket"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	websocket2 "golang.org/x/net/websocket"
	"time"
)

var HttpSet = wire.NewSet(
	ProvideHttpProvider,
	ProvideHttpServerProvider,
)

type HttpProvider struct{}

func ProvideHttpProvider(
	_ HttpServerProvider,
) HttpProvider {
	return HttpProvider{}
}

type HttpServerProvider struct{}

func ProvideHttpServerProvider(
	cfg *config.Config,
	handler *websocket.Handler,
) (HttpServerProvider, error) {
	server := echo.New()

	server.GET("/ws", func(c echo.Context) error {
		websocket2.Handler(handler.Handle).ServeHTTP(c.Response(), c.Request())

		return nil
	})

	errCh := make(chan error, 1)

	go func() {
		if err := server.Start(cfg.Server.Http.Addr); err != nil {
			errCh <- fmt.Errorf("error starting http server. %w", err)
		}
	}()

	select {
	case err := <-errCh:
		return HttpServerProvider{}, fmt.Errorf("error from errCh. %w", err)
	case <-time.After(300 * time.Millisecond):
	}

	return HttpServerProvider{}, nil
}
