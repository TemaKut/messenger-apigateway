package delegatesvc

import (
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
	delegatedto "github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
)

func decodeUserRegisterRequest(req delegatedto.UserRegisterRequest) authdto.RegisterUserRequest {
	return authdto.RegisterUserRequest{
		Name:     req.Name,
		LastName: req.LastName,
		Email:    req.Email,
		Password: req.Password,
	}
}
