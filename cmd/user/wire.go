// +build wireinject

package main

import (
	"caas-micro/internal/app/user"
	"caas-micro/internal/app/user/model"
	"caas-micro/internal/app/user/model/impl/gorm"

	//"caas-micro/internal/app/user/model/impl/gorm"
	//"caas-micro/internal/app/user/model/impl/gorm/entity"
	"caas-micro/internal/app/user/rpcclients"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	user.ProviderSet,
	rpcclients.ProviderSet,
)

func CreateUserServer() (*user.UserServer, error) {
	//panic(wire.Build(providerSet, model.ProviderProductionSet, entity.ProviderSet, gorm.ProviderSet))
	panic(wire.Build(providerSet, model.ProviderProductionSet, gorm.ProviderSet))
}
