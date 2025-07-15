package websocket

import (
	delegatedto "github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
)

func decodeUserRegisterRequest(req *pb.UserRegisterRequest) delegatedto.UserRegisterRequest {
	return delegatedto.UserRegisterRequest{
		Name:     req.GetName(),
		LastName: req.GetLastName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}
