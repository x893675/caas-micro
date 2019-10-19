package main

import (
	"caas-micro/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

var (
	// SERVICENAME microservice name
	SERVICENAME = util.GetEnvironment("SERVICE_NAME", "go.micro.web.caas-micro")

	// WEBADDR web service listen address
	WEBADDR = util.GetEnvironment("WEB_LISTEN_ADDR", "0.0.0.0:8080")
)

func main() {

	// Create service
	service := web.NewService(
		web.Name(SERVICENAME),
	)

	service.Init(
		web.Address(WEBADDR),
	)

	//apiApp := CreateApiApplication()
	// router := gin.Default()
	// router.GET("/v1/greeter", apiApp.LoginCtl.Anything)
	// router.GET("/v1/greeter/:name", apiApp.LoginCtl.Hello)
	//service.Handle("/", router)

	service.Handle("/", InitWeb())

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitWeb() *gin.Engine {
	app := gin.New()

	api := CreateApiApplication()

	api.RegisterRouter(app)

	return app
}
