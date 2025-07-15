package websocket

import (
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
	pbuser "github.com/TemaKut/messenger-client-proto/gen/go/user"
)

func encodeUser(u authdto.User) *pbuser.User {
	return &pbuser.User{
		Id:       u.Id,
		Name:     u.Name,
		LastName: u.LastName,
	}
}
