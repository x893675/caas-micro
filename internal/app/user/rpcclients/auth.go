package rpcclients

import (
	"caas-micro/proto/auth"
	"github.com/micro/go-micro/client"
)

func NewAuthSrvClient() (auth.AuthService, error) {
	return auth.NewAuthService("go.micro.srv.auth", client.DefaultClient), nil
}
