package delegatesvc

import (
	"context"
	"errors"
	"fmt"
	"github.com/TemaKut/messenger-apigateway/internal/app/adapter/auth"
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
		return delegatedto.UserRegisterResponse{}, fmt.Errorf("error register user. %w", encodeError(err))
	}

	return encodeUserRegisterResponse(userRegisterResponse), nil
}

func encodeError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, auth.ErrUserEmailAlreadyExists):
		return fmt.Errorf("%w, %w", delegatedto.ErrUserEmailAlreadyExists, err)
	default:
		return fmt.Errorf("%w, %w", delegatedto.ErrUnknown, err)
	}
}
