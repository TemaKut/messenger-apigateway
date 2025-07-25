package delegatedto

type UserRegisterRequest struct {
	Name     string
	LastName string
	Email    string
	Password string
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
