package controllers

import (
	"fmt"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
)

type UserRegisterController interface {
	Invoke(req Request) error
}

type UserRegisterControllerImpl struct {
}

func NewUserRegisterControllerImpl() *UserRegisterControllerImpl {
	return &UserRegisterControllerImpl{}
}

func (i *UserRegisterControllerImpl) Invoke(req Request) error {
	err := req.SendResponse(req.Ctx(), &pb.Response{
		Source: &pb.Response_Success{
			Success: &pb.Success{
				Data: &pb.Success_Auth{
					Auth: &pb.AuthResponse{
						Username: "Hello",
						Password: "123456",
					},
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("error send response. %w", err)
	}

	return nil
}
