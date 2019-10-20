package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro"

	"context"

	"caas-micro/proto/auth"
	"caas-micro/proto/user"
)

type Auth struct {
	userSVc user.UserService
}

func (a *Auth) GenerateToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Println("in GenerateToken")
	log.Println(req.Username)
	log.Println(req.Password)
	rsp.Msg = "Hello " + req.Username
	return nil
}

func (a *Auth) DestroyToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Println("in DestroyToken")
	log.Println(req.Username)
	log.Println(req.Password)
	rsp.Msg = "Hello " + req.Username
	return nil
}

func (a *Auth) Verify(ctx context.Context, req *auth.LoginRequest, rsp *auth.Token) error {
	log.Println("in Verify")
	response, err := a.userSVc.Query(ctx, &user.Request{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	rsp.AccessToken = response.Msg
	rsp.TokenType = "test"
	rsp.ExpiresAt = 2000000
	return nil
}

func (a *Auth) VertifyToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Println("in VertifyToken")
	log.Println(req.Username)
	log.Println(req.Password)
	if req.Username == "test" && req.Password == "test" {
		rsp.Msg = "userid"
		return nil
	}
	return errors.New("invalid token")
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()

	// authServer := &Auth{
	// 	userSVc: user.NewUserService("go.micro.srv.user", client.DefaultClient),
	// }

	authServer, err := CreateAuthServer()
	if err != nil {
		log.Fatal(err)
	}

	auth.RegisterAuthHandler(service.Server(), authServer)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
