package rpcclients

import (
	"caas-micro/proto/auth"

	"github.com/micro/go-micro/client"
)

func NewAuthSrvClient() auth.AuthService {
	return auth.NewAuthService("go.micro.srv.auth", client.DefaultClient)
}
