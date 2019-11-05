package api

import (
	"caas-micro/internal/app/api/controller"

	"github.com/google/wire"
)

type ApiApplication struct {
	LoginCtl *controller.LoginController
	UserCtl  *controller.UserController
}

func NewApiApplication(loginctl *controller.LoginController, userctl *controller.UserController) *ApiApplication {
	return &ApiApplication{
		LoginCtl: loginctl,
		UserCtl:  userctl,
	}
}

var ProviderSet = wire.NewSet(NewApiApplication)
