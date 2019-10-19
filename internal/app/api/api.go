package api

import (
	"caas-micro/internal/app/api/controller"

	"github.com/google/wire"
)

type ApiApplication struct {
	LoginCtl *controller.LoginController
}

func NewApiApplication(loginctl *controller.LoginController) *ApiApplication {
	return &ApiApplication{
		LoginCtl: loginctl,
	}
}

var ProviderSet = wire.NewSet(NewApiApplication)
