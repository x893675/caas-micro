package proto

import (
	"caas-micro/proto/auth"

	"github.com/micro/go-micro/client"
)

// func Inject(container *dig.Container) error {

// 	return nil
// }

func InitAuth() auth.AuthService {
	cl := auth.NewAuthService("go.micro.srv.auth", client.DefaultClient)
	return cl
}
