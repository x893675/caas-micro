package api

import "github.com/gin-gonic/gin"

func (api *ApiApplication) RegisterRouter(app *gin.Engine) {
	g := app.Group("/api")
	v1 := g.Group("/v1")
	{
		v1.GET("/greeter", api.LoginCtl.Anything)
		v1.GET("/greeter/:name", api.LoginCtl.Hello)
		v1.POST("/login", api.LoginCtl.Login)
	}
}
