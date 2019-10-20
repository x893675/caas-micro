package main

import (
	"caas-micro/proto/user"
	"context"
	"log"
	"time"

	"github.com/micro/go-micro"
)

type User struct{}

func (user *User) Query(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Println("in usersvc.Query")
	log.Println(req.Username)
	log.Println(req.Password)
	rsp.Msg = "login sucessful"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	//hello.RegisterSayHandler(service.Server(), new(Say))
	user.RegisterUserHandler(service.Server(), new(User))
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
