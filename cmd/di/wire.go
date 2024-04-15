package di

import (
	"github.com/google/wire"
	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/apimanager"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/s-yakubovskiy/whereami/internal/dumper"
	"github.com/s-yakubovskiy/whereami/internal/whereami"
	"github.com/s-yakubovskiy/whereami/pkg/gpsdfetcher"
	"github.com/s-yakubovskiy/whereami/pkg/ipconfig"
	"github.com/spf13/cobra"
)

// This function is just a placeholder to create a provider set
func InitializeStoreCommand() (*cobra.Command, error) {
	wire.Build(
		config.Cfg,
		ipconfig.NewIPConfig,
		apimanager.NewAPIManager,
		dbclient.NewSQLiteDB,
		dumper.NewDumperJSON,
		gpsdfetcher.NewGPSDFetcher,
		whereami.NewLocator,
		wire.Struct(new(whereami.Config), "*"), // Assuming Config struct that holds all configurations
		wire.Struct(new(cobra.Command), "*"),
	)
	return nil, nil
}
