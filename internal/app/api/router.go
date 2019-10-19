package api

import "github.com/gin-gonic/gin"

func (api *ApiApplication) RegisterRouter(app *gin.Engine) {
	g := app.Group("/api")
	v1 := g.Group("/v1")
	{
		v1.GET("/v1/greeter", api.LoginCtl.Anything)
		v1.GET("/v1/greeter/:name", api.LoginCtl.Hello)
	}
}
