package websocket

import (
	delegatedto "github.com/TemaKut/messenger-apigateway/internal/dto/delegate"
	pb "github.com/TemaKut/messenger-client-proto/gen/go"
	pbuser "github.com/TemaKut/messenger-client-proto/gen/go/user"
	"google.golang.org/protobuf/types/known/timestamppb"
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
				User:       encodeUser(resp.User),
				AuthParams: encodeAuthParams(resp.AuthParams),
			},
		},
	}
}

func encodeAuthParams(params delegatedto.AuthParams) *pbuser.AuthParams {
	return &pbuser.AuthParams{
		AccessToken:  encodeAuthToken(params.AccessToken),
		RefreshToken: encodeAuthToken(params.RefreshToken),
	}
}

func encodeAuthToken(token delegatedto.AuthToken) *pbuser.AuthToken {
	return &pbuser.AuthToken{
		Token:     token.Token,
		ExpiredAt: timestamppb.New(token.ExpiredAt),
	}
}

func encodeUser(u delegatedto.User) *pbuser.User {
	return &pbuser.User{
		Id:       u.Id,
		Name:     u.Name,
		LastName: u.LastName,
	}
}
