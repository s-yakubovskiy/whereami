package dbclient

import (
	"github.com/google/wire"
	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

func ProvideDatabase(cfg config.Database) (contracts.LKInterface, error) {
	return NewSQLiteDB(cfg)
}

var ProviderSet = wire.NewSet(
	ProvideDatabase,
)
