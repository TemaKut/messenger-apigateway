package websocket

import (
	delegatedto "github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
	pbuser "github.com/TemaKut/messenger-client-proto/gen/go/user"
)

func encodeUserRegisterResponse(resp delegatedto.UserRegisterResponse) *pb.Success {
	return &pb.Success{
		Data: &pb.Success_UserRegister{
			UserRegister: &pb.UserRegisterResponse{
				User: encodeUser(resp.User),
			},
		},
	}
}

func encodeUserAuthorizeResponse(resp delegatedto.UserAuthorizeResponse) *pb.Success {
	return &pb.Success{
		Data: &pb.Success_UserAuthorize{
			UserAuthorize: &pb.UserAuthorizeResponse{
				//: encodeUser(resp.User), // TODO
			},
		},
	}
}

func encodeUser(u delegatedto.User) *pbuser.User {
	return &pbuser.User{
		Id:       u.Id,
		Name:     u.Name,
		LastName: u.LastName,
	}
}
