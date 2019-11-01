package main

import (
	"caas-micro/internal/app/api/middleware"
	"caas-micro/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

var (
	// SERVICENAME microservice name
	SERVICENAME = util.GetEnvironment("SERVICE_NAME", "go.micro.web.caas-micro.api")

	// WEBADDR web service listen address
	WEBADDR = util.GetEnvironment("WEB_LISTEN_ADDR", "0.0.0.0:8080")

	RUNMODE = util.GetEnvironment("RUN_MODE", "debug")
)

func main() {

	// Create service
	service := web.NewService(
		web.Name(SERVICENAME),
	)

	_ = service.Init(
		web.Address(WEBADDR),
	)

	service.Handle("/", InitWeb())

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitWeb() *gin.Engine {
	gin.SetMode(RUNMODE)

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())
	// 崩溃恢复
	app.Use(middleware.RecoveryMiddleware())
	api := CreateApiApplication()

	api.RegisterRouter(app)

	return app
}
