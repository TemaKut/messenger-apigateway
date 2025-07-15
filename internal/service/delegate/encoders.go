package delegatesvc

import (
	authdto "github.com/TemaKut/messenger-apigateway/internal/dto/auth"
	delegatedto "github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
)

func encodeUserRegisterResponse(response authdto.RegisterUserResponse) delegatedto.UserRegisterResponse {
	return delegatedto.UserRegisterResponse{
		User: encodeUser(response.User),
	}
}

func encodeUser(user authdto.User) delegatedto.User {
	return delegatedto.User{
		Id:       user.Id,
		Name:     user.Name,
		LastName: user.LastName,
	}
}
