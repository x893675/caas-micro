package rpcclients

import (
	"caas-micro/proto/user"
	"github.com/micro/go-micro/client"
)

func NewUserSrvClient() user.UserService {
	return user.NewUserService("go.micro.srv.user", client.DefaultClient)
}
