package controllers

type Controller struct {
	Auth struct {
		UserRegisterController UserRegisterController
	}
}

func NewController(
	userRegisterController UserRegisterController,
) *Controller {
	var controller Controller

	controller.Auth.UserRegisterController = userRegisterController

	return &controller
}
