package main

import (
	"caas-micro/proto/user"
	"log"
	"time"

	"github.com/micro/go-micro"
)

//type User struct{}
//
//func (user *User) Query(ctx context.Context, req *user.QueryRequest, rsp *user.QueryResult) error {
//	log.Println("in usersvc.Query")
//	log.Println(req.UserName)
//	//log.Println(req.)
//	return nil
//}

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
	//user.RegisterUserHandler(service.Server(), new(User))

	userServer, err := CreateUserServer()
	if err != nil {
		log.Fatal(err)
	}
	_ = user.RegisterUserHandler(service.Server(), userServer)
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
