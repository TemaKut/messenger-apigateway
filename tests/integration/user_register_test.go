package integration

import (
	"context"
	"fmt"
	"github.com/TemaKut/messenger-apigateway/tests/integration/client"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
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
		Id: 2,
		Data: &pb.Request_UserRegister{
			UserRegister: &pb.UserRegisterRequest{
				Name:     "Name",
				LastName: "LastName",
				Email:    "email@email.ru",
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
