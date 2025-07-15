package delegatesvc

import (
	"context"
	"fmt"
	delegatedto "github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
)

// Service делегирует запросы в другие сервисы. + допускает небольшую бизнес-логику при необходимости
// Для некоторых GET запросов стоит реализовать кеширование при помощи слоя storage
// Именование методов On[название_реквеста_клиентского_прото]
// Метод принимает (ctx, dto клиентского запроса [...Request])
// Метод возвращает опционально dto [...Response] и опционально error
type Service struct {
	authService AuthService
}

func NewService(authService AuthService) *Service {
	return &Service{authService: authService}
}

func (s *Service) OnUserRegisterRequest(
	ctx context.Context,
	req delegatedto.UserRegisterRequest,
) (delegatedto.UserRegisterResponse, error) {
	userRegisterResponse, err := s.authService.RegisterUser(ctx, decodeUserRegisterRequest(req))
	if err != nil {
		return delegatedto.UserRegisterResponse{}, fmt.Errorf("error register user. %w", err)
	}

	return encodeUserRegisterResponse(userRegisterResponse), nil
}
