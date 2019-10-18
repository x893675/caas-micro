package main

import (
	"caas-micro/internal/app/api"
	"caas-micro/internal/app/api/controller"
	"caas-micro/proto/auth"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
)

var (
	authSrcCl auth.AuthService
)

func main() {

	// Create service
	service := web.NewService(
		web.Name("go.micro.api.caas-micro"),
	)

	service.Init()

	// setup Greeter Server Client
	authSrcCl = auth.NewAuthService("go.micro.srv.auth", client.DefaultClient)
	loginctl := controller.NewLoginController(authSrcCl)
	apiApp := api.NewApiApplication(loginctl)

	// Create RESTful handler (using Gin)
	//say := new(Say)
	router := gin.Default()
	router.GET("/greeter", apiApp.LoginCtl.Anything)
	router.GET("/greeter/:name", apiApp.LoginCtl.Hello)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
