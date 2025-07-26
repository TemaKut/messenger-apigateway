package authdto

import "time"

type RegisterUserRequest struct {
	Name     string
	LastName string
	Email    string
	Password string
}

type RegisterUserResponse struct {
	User User
}

type UserAuthorizeRequest struct {
	Credentials UserAuthorizeCredentials
}

type UserAuthorizeCredentials struct {
	Email *UserAuthorizeEmailCredential
}

type UserAuthorizeEmailCredential struct {
	Email    string
	Password string
}

type UserAuthorizeResponse struct {
	User       User
	AuthParams AuthParams
}

type AuthParams struct {
	AccessToken  AuthToken
	RefreshToken AuthToken
}

type AuthToken struct {
	Token     string
	ExpiredAt time.Time
}
