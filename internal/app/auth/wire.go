package auth

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewAuther, NewAuthServer)
