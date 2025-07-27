package userregister

import (
	"fmt"
	"github.com/TemaKut/messenger-apigateway/tests/integration/client"
	"github.com/TemaKut/messenger-apigateway/tests/integration/config"
	"github.com/TemaKut/messenger-apigateway/tests/integration/utils"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	client *client.Client
}

func (s *Suite) SetupSuite() {
	cfg := config.NewTestConfig()

	cl, err := client.NewClient(s.T().Context(), cfg.TestApiGatewayWsAddr)
	if err != nil {
		s.T().Fatalf("error make client. %s", err)
	}

	s.client = cl
}

func (s *Suite) TestUserRegister() {
	tt
	registerResponse, err := s.client.Request(s.T().Context(), &pb.Request{
		Id: uuid.New().String(),
		Data: &pb.Request_UserRegister{
			UserRegister: &pb.UserRegisterRequest{
				Name:     "TestName",
				LastName: "TestLastName",
				Email:    utils.RandomUserEmail(),
				Password: "123456",
			},
		},
	})
	if err != nil {
		s.T().Fatal(fmt.Errorf("error user register request. %w", err))
	}

	fmt.Println(registerResponse)
}

func (s *Suite) TearDownSuite() {
	s.client.Close()
}
