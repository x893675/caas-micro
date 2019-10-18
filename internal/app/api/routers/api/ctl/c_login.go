package ctl

import (
	"caas-micro/proto/auth"

	"github.com/gin-gonic/gin"
)

// NewLogin 创建登录管理控制器
func NewLogin(authSrvCl auth.AuthService) *Login {
	return &Login{
		AtuhSrvClient: authSrvCl,
	}
}

type Login struct {
	AtuhSrvClient auth.AuthService
}

func (a *Login) Anything(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}
