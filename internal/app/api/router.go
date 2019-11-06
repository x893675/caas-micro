package api

import (
	"caas-micro/internal/app/api/middleware"

	"github.com/gin-gonic/gin"
)

func (api *ApiApplication) RegisterRouter(app *gin.Engine) {
	g := app.Group("/api")
	// 用户身份授权
	g.Use(middleware.UserAuthMiddleware(
		api.LoginCtl.AuthSvc,
		middleware.AllowMethodAndPathPrefixSkipper(
			middleware.JoinRouter("POST", "/api/v1/login"),
			//middleware.JoinRouter("GET", "/api/v1/users"),
		),
	))
	v1 := g.Group("/v1")
	{
		//v1.GET("/greeter", api.LoginCtl.Anything)
		//v1.GET("/greeter/:name", api.LoginCtl.Hello)
		v1.POST("/login", api.LoginCtl.Login)

		// 注册/api/v1/current
		v1.PUT("/current/password", api.UserCtl.GetUserInfo)

		v1.GET("/users", api.UserCtl.Query)
		v1.GET("/users/:id", api.UserCtl.Get)
		v1.POST("/users", api.UserCtl.Create)
		v1.PUT("/users/:id", api.UserCtl.Update)
		v1.DELETE("/users/:id", api.UserCtl.Delete)
		v1.PATCH("/users/:id/enable", api.UserCtl.Enable)
		v1.PATCH("/users/:id/disable", api.UserCtl.Disable)
	}
}
