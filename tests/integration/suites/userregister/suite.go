package userregister

import (
	"fmt"
	"github.com/TemaKut/messenger-apigateway/tests/integration/client"
	"github.com/TemaKut/messenger-apigateway/tests/integration/config"
	"github.com/TemaKut/messenger-apigateway/tests/integration/utils"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"strings"
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
	tt := []struct {
		name                string
		userRegisterRequest *pb.UserRegisterRequest
	}{
		{
			name: "register user successfully",
			userRegisterRequest: &pb.UserRegisterRequest{
				Name:     "TestName",
				LastName: "TestLastName",
				Email:    utils.RandomUserEmail(),
				Password: "12312312",
			},
		},
	}

	for _, test := range tt {
		s.Run(test.name, func() {
			response, err := s.client.Request(s.T().Context(), &pb.Request{
				Id: uuid.New().String(),
				Data: &pb.Request_UserRegister{
					UserRegister: test.userRegisterRequest,
				},
			})
			if err != nil {
				s.T().Fatal(fmt.Errorf("error user register userRegisterRequest. %w", err))
			}

			userRegisterResponse := response.GetSuccess().GetUserRegister()

			if !s.Assert().NotNil(userRegisterResponse, "response is nil") {
				return
			}

			if !s.Assert().Equal(userRegisterResponse.GetUser().GetName(), test.userRegisterRequest.GetName()) {
				return
			}

			if !s.Assert().Equal(userRegisterResponse.GetUser().GetLastName(), test.userRegisterRequest.GetLastName()) {
				return
			}

			if s.Assert().NoError(uuid.Validate(userRegisterResponse.GetUser().GetId())) {
				return
			}
		})
	}
}

func (s *Suite) TestUserRegisterWithNotUniqueEmail() {
	email := utils.RandomUserEmail()

	response, err := s.client.Request(s.T().Context(), &pb.Request{
		Id: uuid.New().String(),
		Data: &pb.Request_UserRegister{
			UserRegister: &pb.UserRegisterRequest{
				Name:     "TestName",
				LastName: "TestLastName",
				Email:    email,
				Password: "12312312",
			},
		},
	})
	if err != nil {
		s.T().Fatal(fmt.Errorf("error user register userRegisterRequest. %w", err))
	}

	if !s.Assert().NotNil(response.GetSuccess().GetUserRegister()) {
		return
	}

	responseSecond, err := s.client.Request(s.T().Context(), &pb.Request{
		Id: uuid.New().String(),
		Data: &pb.Request_UserRegister{
			UserRegister: &pb.UserRegisterRequest{
				Name:     "TestName2",
				LastName: "TestLastName2",
				Email:    email,
				Password: "12312312",
			},
		},
	})
	if err != nil {
		s.T().Fatal(fmt.Errorf("error user register userRegisterRequest. %w", err))
	}

	if !s.Assert().Equal(pb.ErrorReason_ERROR_REASON_USER_EMAIL_ALREADY_EXISTS, responseSecond.GetError().GetReason()) {
		return
	}
}

func (s *Suite) TestUserRegisterValidationErrors() {
	tt := []struct {
		name                string
		userRegisterRequest *pb.UserRegisterRequest
		errorReason         pb.ErrorReason
	}{
		{
			name: "without name",
			userRegisterRequest: &pb.UserRegisterRequest{
				Name:     "",
				LastName: "TestLastName",
				Email:    utils.RandomUserEmail(),
				Password: "12312312",
			},
			errorReason: pb.ErrorReason_ERROR_REASON_VALIDATION,
		},
		{
			name: "without email",
			userRegisterRequest: &pb.UserRegisterRequest{
				Name:     "TestName2",
				LastName: "TestLastName",
				Email:    "",
				Password: "12312312",
			},
			errorReason: pb.ErrorReason_ERROR_REASON_VALIDATION,
		},
		{
			name: "with invalid email",
			userRegisterRequest: &pb.UserRegisterRequest{
				Name:     "TestName2",
				LastName: "TestLastName",
				Email:    "dsfsdfsddssdfdssderge",
				Password: "12312312",
			},
			errorReason: pb.ErrorReason_ERROR_REASON_VALIDATION,
		},
		{
			name: "without password",
			userRegisterRequest: &pb.UserRegisterRequest{
				Name:     "TestName2",
				LastName: "TestLastName",
				Email:    utils.RandomUserEmail(),
				Password: "",
			},
			errorReason: pb.ErrorReason_ERROR_REASON_VALIDATION,
		},
		{
			name: "short password",
			userRegisterRequest: &pb.UserRegisterRequest{
				Name:     "TestName2",
				LastName: "TestLastName",
				Email:    utils.RandomUserEmail(),
				Password: "1234567",
			},
			errorReason: pb.ErrorReason_ERROR_REASON_VALIDATION,
		},
		{
			name: "too long password",
			userRegisterRequest: &pb.UserRegisterRequest{
				Name:     "TestName2",
				LastName: "TestLastName",
				Email:    utils.RandomUserEmail(),
				Password: strings.Repeat("a", 41),
			},
			errorReason: pb.ErrorReason_ERROR_REASON_VALIDATION,
		},
	}

	for _, test := range tt {
		s.Run(test.name, func() {
			response, err := s.client.Request(s.T().Context(), &pb.Request{
				Id: uuid.New().String(),
				Data: &pb.Request_UserRegister{
					UserRegister: test.userRegisterRequest,
				},
			})
			if !s.Assert().NoError(err) {
				return
			}

			if !s.Assert().Equal(test.errorReason, response.GetError().GetReason()) {
				return
			}
		})
	}
}

func (s *Suite) TearDownSuite() {
	s.client.Close()
}
