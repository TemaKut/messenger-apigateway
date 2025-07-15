package authdto

type RegisterUserRequest struct {
	Name     string
	LastName string
	Email    string
	Password string
}

type RegisterUserResponse struct {
	User User
}
