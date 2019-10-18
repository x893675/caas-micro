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
			middleware.JoinRouter("GET", "/api/v1/login"),
			middleware.JoinRouter("POST", "/api/v1/login"),
			middleware.JoinRouter("GET", "/api/v1/auth"),
			//middleware.JoinRouter("GET", "/api/v1/users"),
		),
	))
	v1 := g.Group("/v1")
	{
		v1.POST("/login", api.LoginCtl.Login)
		v1.POST("/login/exit", api.LoginCtl.Logout)
		v1.GET("/auth", api.LoginCtl.OpenshiftLogin)

		// 注册/api/v1/current
		v1.GET("/current/user", api.UserCtl.GetUserInfo)
		v1.PUT("/current/password", api.UserCtl.UpdatePassword)

		v1.GET("/users", api.UserCtl.Query)
		v1.GET("/users/:id", api.UserCtl.Get)
		v1.POST("/users", api.UserCtl.Create)
		v1.PUT("/users/:id", api.UserCtl.Update)
		v1.DELETE("/users/:id", api.UserCtl.Delete)
		v1.PATCH("/users/:id/enable", api.UserCtl.Enable)
		v1.PATCH("/users/:id/disable", api.UserCtl.Disable)
	}
}
