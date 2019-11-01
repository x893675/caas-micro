package gorm

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewDB)
