package factory

import (
	"fmt"
	"github.com/TemaKut/messenger-apigateway/internal/app/config"
	"github.com/TemaKut/messenger-apigateway/internal/app/logger"
	authv1 "github.com/TemaKut/messenger-service-proto/gen/go/auth"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var GrpcClientsSet = wire.NewSet(
	ProvideUserApiGrpcClient,
)

func ProvideUserApiGrpcClient(
	cfg *config.Config,
	logger *logger.Logger,
) (authv1.UserAPIClient, func(), error) {
	logger.Infof("connecting user api grpc client on: %s", cfg.Clients.UserApi.Addr)

	conn, err := grpc.NewClient(cfg.Clients.UserApi.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("errr make grpc client. %w", err)
	}

	return authv1.NewUserAPIClient(conn), func() {
		logger.Infof("close user api grpc client")

		if err := conn.Close(); err != nil {
			logger.Errorf("close grpc client connection error. %s", err)
		}
	}, nil
}
