package integration

import (
	"context"
	"fmt"
	"github.com/TemaKut/messenger-apigateway/tests/integration/client"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestUserRegister(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cli, err := client.NewClient(ctx, "ws://localhost:8000/ws")
	if err != nil {
		t.Fatalf("error make client. %s", err)

		return
	}

	defer cli.Close()

	response, err := cli.Request(context.Background(), &pb.Request{
		Id: uuid.NewString(),
		Data: &pb.Request_UserRegister{
			UserRegister: &pb.UserRegisterRequest{
				Name:     "Name3",
				LastName: "LastName3",
				Email:    "email3@email.ru",
				Password: "123123",
			},
		},
	})
	if err != nil {
		t.Fatalf("error request. %s", err)
		return
	}

	fmt.Println(response)
}

func TestUserAuthorizeByEmail(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cli, err := client.NewClient(ctx, "ws://localhost:8000/ws")
	if err != nil {
		t.Fatalf("error make client. %s", err)

		return
	}

	defer cli.Close()

	response, err := cli.Request(context.Background(), &pb.Request{
		Id: uuid.NewString(),
		Data: &pb.Request_UserAuthorize{
			UserAuthorize: &pb.UserAuthorizeRequest{
				Credentials: &pb.UserAuthorizeRequest_Email{
					Email: &pb.UserAuthorizeEmailCredential{
						Email:    "email3@email.ru",
						Password: "123123",
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("error request. %s", err)
		return
	}

	fmt.Println(response)
}
