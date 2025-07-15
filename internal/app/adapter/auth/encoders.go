package auth

import (
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
	authv1 "github.com/TemaKut/messenger-service-proto/gen/go/auth"
)

func encodeUser(user *authv1.User) authdto.User {
	return authdto.User{
		Id:       user.GetId(),
		Name:     user.GetName(),
		LastName: user.GetLastName(),
	}
}
