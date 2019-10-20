// +build wireinject

package main

import (
	"caas-micro/internal/app/auth"
	"caas-micro/internal/app/auth/rpcclients"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	auth.ProviderSet,
	rpcclients.ProviderSet,
)

func CreateAuthServer() (*auth.AuthServer, error) {
	panic(wire.Build(providerSet))
}
