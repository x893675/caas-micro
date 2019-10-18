package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"

	"context"

	"caas-micro/proto/auth"
)

type Auth struct{}

func (a *Auth) GenerateToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Println("in GenerateToken")
	log.Println(req.Username)
	log.Println(req.Password)
	rsp.Msg = "Hello " + req.Username
	return nil
}

func (a *Auth) VertifyToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Println("in VertifyToken")
	log.Println(req.Username)
	log.Println(req.Password)
	rsp.Msg = "Hello " + req.Username
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	//hello.RegisterSayHandler(service.Server(), new(Say))
	auth.RegisterAuthHandler(service.Server(), new(Auth))
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
