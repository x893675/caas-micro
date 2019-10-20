package controller

import (
	"caas-micro/internal/app/api/pkg/ginplus"
	"caas-micro/proto/auth"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	authSvc auth.AuthService
}

func NewLoginController(auth auth.AuthService) *LoginController {
	return &LoginController{
		authSvc: auth,
	}
}

func (s *LoginController) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (s *LoginController) Hello(c *gin.Context) {
	log.Print("Received Say.Hello API request")

	name := c.Param("name")

	response, err := s.authSvc.GenerateToken(context.TODO(), &auth.Request{
		Username: name,
		Password: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}

func (s *LoginController) Login(c *gin.Context) {
	log.Print("Received api.Login request")

	var item auth.LoginRequest
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	response, err := s.authSvc.Verify(context.TODO(), &item)

	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, response)
	// name := c.Param("name")

	// response, err := s.authSvc.GenerateToken(context.TODO(), &auth.Request{
	// 	Username: name,
	// 	Password: name,
	// })
}
