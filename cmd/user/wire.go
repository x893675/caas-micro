// +build wireinject

package main

import (
	"caas-micro/internal/app/user"
	"caas-micro/internal/app/user/model"
	"caas-micro/internal/app/user/pkg/gormplus"
	"caas-micro/internal/app/user/rpcclients"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	user.ProviderSet,
	rpcclients.ProviderSet,
)

func CreateUserServer() (*user.UserServer, error) {
	panic(wire.Build(providerSet, gormplus.ProviderSet, model.ProviderProductionSet))
}
