package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"caas-micro/internal/app/api/routers/api/ctl"
)

// RegisterRouter 注册/api路由
func RegisterRouter(app *gin.Engine, container *dig.Container) error {
	err := ctl.Inject(container)
	if err != nil {
		return err
	}

	return container.Invoke(func(login *ctl.Login) error {
		g := app.Group("/api")

		v1 := g.Group("/v1")
		{
			v1.GET("/hello", login.Anything)
		}
		return nil
	})
}
