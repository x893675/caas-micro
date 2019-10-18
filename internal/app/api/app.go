package api

import "caas-micro/internal/app/api/controller"

type ApiApplication struct {
	LoginCtl *controller.LoginController
}

func NewApiApplication(loginctl *controller.LoginController) *ApiApplication {
	return &ApiApplication{
		LoginCtl: loginctl,
	}
}
