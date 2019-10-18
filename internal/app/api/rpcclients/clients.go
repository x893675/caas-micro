package rpcclients

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewAuthSrvClient, NewUserSrvClient)
