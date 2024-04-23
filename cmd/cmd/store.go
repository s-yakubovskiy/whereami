package cmd

import (
	"log"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/apimanager"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/s-yakubovskiy/whereami/internal/dumper"
	"github.com/s-yakubovskiy/whereami/internal/servicefactory"
	"github.com/s-yakubovskiy/whereami/internal/whereami"
	"github.com/s-yakubovskiy/whereami/pkg/gpsdfetcher"
	"github.com/s-yakubovskiy/whereami/pkg/ipconfig"
	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store WhereAmI application",
	Long:  `This command stores location information in the database (sqlite).`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg
		if locationApi != "" {
			cfg.MainProvider = locationApi
		}

		factory := &servicefactory.DefaultServiceFactory{}

		ipconfig, err := ipconfig.NewIPConfig(cfg.ProviderConfigs.PublicIpProvider)
		if err != nil {
			log.Fatalf("Failed to create IP configuration: %v", err)
		}

		ipapi, err := factory.CreateLocationService(cfg.ProviderConfigs.IpApi)
		if err != nil {
			log.Fatalf("Failed to create primary location service: %v", err)
		}
		ipdata, err := factory.CreateLocationService(cfg.ProviderConfigs.IpData)
		if err != nil {
			log.Fatalf("Failed to create secondary location service: %v", err)
		}
		ipquality, err := factory.CreateQualityService(cfg.ProviderConfigs.IpQualityScore)
		if err != nil {
			log.Fatalf("Failed to create IP quality service: %v", err)
		}

		client := apimanager.NewAPIManager(ipconfig, ipapi, ipdata, ipquality)
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		dumper, err := dumper.NewDumperJSON(dbcli)
		if err != nil {
			log.Fatalf("Failed to create dumper: %v", err)
		}

		var gps gpsdfetcher.GPSInterface
		if cfg.GPSConfig.Enabled || gpsEnabled {
			cfg.GPSConfig.Enabled = true
			gps = gpsdfetcher.NewGPSDFetcher(cfg.GPSConfig.Timeout)
		}

		lCfg := whereami.NewConfig(cfg.ProviderConfigs.IpQualityScore.Enabled, ipLookup, gpsEnabled)
		locator := whereami.NewLocator(client, dbcli, dumper, gps, lCfg)
		locator.Store()
	},
}

func init() {
	storeCmd.Flags().StringVarP(&locationApi, "provider", "p", "", "Select IP location provider: [ipapi, ipdata]")
	rootCmd.AddCommand(storeCmd)
}
