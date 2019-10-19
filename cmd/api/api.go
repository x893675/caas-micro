package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

// var (
// 	authSrcCl auth.AuthService
// )

func main() {

	// Create service
	service := web.NewService(
		web.Name("go.micro.web.caas-micro.api"),
		web.Address("0.0.0.0:8080"),
	)

	service.Init()

	// setup Greeter Server Client
	// authSrcCl = auth.NewAuthService("go.micro.srv.auth", client.DefaultClient)
	// loginctl := controller.NewLoginController(authSrcCl)
	// apiApp := api.NewApiApplication(loginctl)
	apiApp := CreateApiApplication()
	// Create RESTful handler (using Gin)
	//say := new(Say)
	router := gin.Default()
	router.GET("/v1/greeter", apiApp.LoginCtl.Anything)
	router.GET("/v1/greeter/:name", apiApp.LoginCtl.Hello)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
