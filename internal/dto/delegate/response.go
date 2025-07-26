package delegatedto

import "time"

type UserRegisterResponse struct {
	User User
}

type User struct {
	Id       string
	Name     string
	LastName string
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
