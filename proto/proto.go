package proto

import (
	"caas-micro/proto/auth"

	"github.com/micro/go-micro/client"
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	container.Provide(auth.NewAuthService("go.micro.srv.auth", client.DefaultClient))
	return nil
}
