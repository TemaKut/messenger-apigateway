package delegatedto

type UserRegisterResponse struct {
	User User
}

type User struct {
	Id       string
	Name     string
	LastName string
}
