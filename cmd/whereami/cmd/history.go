//go:build ignore
// +build ignore

package cmd

import (
	"log"

	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/apimanager"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/s-yakubovskiy/whereami/internal/dumper"
	"github.com/s-yakubovskiy/whereami/internal/servicefactory"
	"github.com/s-yakubovskiy/whereami/internal/whereami"
	"github.com/s-yakubovskiy/whereami/pkg/gpsdfetcher"
	"github.com/spf13/cobra"
)

var NumHistory int32

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "History of WhereAmI locations",
	Long:  `This command outputs the history of whereami locations. Print to stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		introduce()
		cfg := config.Cfg
		if locationApi != "" {
			cfg.MainProvider = locationApi
		}

		factory := &servicefactory.DefaultServiceFactory{}

		ifconfig, err := factory.CreateIpProviderService(cfg.ProviderConfigs.Ifconfig)
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

		locationService := apimanager.NewFallbackLocationService(ipapi, ipdata)
		client := apimanager.NewAPIManager(ifconfig, locationService, ipquality)
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

		lCfg := whereami.NewConfig(cfg.ProviderConfigs.IpQualityScore.Enabled, publicIP, gpsEnabled)
		locator := whereami.NewLocator(client, dbcli, dumper, gps, lCfg)
		locator.History(NumHistory)
	},
}

func init() {
	historyCmd.Flags().Int32VarP(&NumHistory, "number", "n", 10, "Specify number of last history locations to output")
	rootCmd.AddCommand(historyCmd)
}
