// +build wireinject

package main

import (
	"caas-micro/internal/app/api"
	"caas-micro/internal/app/api/controller"
	"caas-micro/internal/app/api/rpcclients"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	api.ProviderSet,
	controller.ProviderSet,
	rpcclients.ProviderSet,
)

func CreateApiApplication() *api.ApiApplication {
	panic(wire.Build(providerSet))
}
