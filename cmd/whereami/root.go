package main

import (
	"fmt"
	"log"
	"os"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/apimanager"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/s-yakubovskiy/whereami/internal/dumper"
	"github.com/s-yakubovskiy/whereami/internal/whereami"
	"github.com/s-yakubovskiy/whereami/pkg/ipconfig"
	"github.com/spf13/cobra"
)

func getLocationClient(provider string) (apimanager.IPLocationInterface, apimanager.IPLocationInterface, apimanager.IPQualityInterface, error) {
	switch provider {
	case "ipapi":
		primary, err := apimanager.NewIpApiClient(config.Cfg.ProviderConfigs.IpApi)
		secondary, err := apimanager.NewIpDataClient(config.Cfg.ProviderConfigs.IpData)
		if config.Cfg.ProviderConfigs.IpQualityScore.Enabled {
			ipquality, err := apimanager.NewIpQualityScoreClient(config.Cfg.ProviderConfigs.IpQualityScore)
			return primary, secondary, ipquality, err
		}
		return primary, secondary, nil, err
	case "ipdata":
		primary, err := apimanager.NewIpDataClient(config.Cfg.ProviderConfigs.IpData)
		secondary, err := apimanager.NewIpApiClient(config.Cfg.ProviderConfigs.IpApi)
		if config.Cfg.ProviderConfigs.IpQualityScore.Enabled {
			ipquality, err := apimanager.NewIpQualityScoreClient(config.Cfg.ProviderConfigs.IpQualityScore)
			return primary, secondary, ipquality, err
		}
		return primary, secondary, nil, err
	default:
		log.Fatalf("Unknown provider set: %+v\n", provider)
		return nil, nil, nil, fmt.Errorf("unknown provider: %s", provider)
	}
}

var rootCmd = &cobra.Command{
	Use:   "whereami",
	Short: "WhereAmI is an application to find your geolocation based on your IP",
	Long:  `WhereAmI is a CLI application that allows users to find their geolocation based on their public IP address.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg

		ipconfig, err := ipconfig.NewIPConfig()
		primary, secondary, ipquality, err := getLocationClient(cfg.MainProvider)

		client := apimanager.NewAPIManager(ipconfig, primary, secondary, ipquality)
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		dumper, err := dumper.NewDumperJSON(dbcli)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open database: %v\n", err)
			os.Exit(1)
		}
		locator := whereami.NewLocator(client, dbcli, dumper, cfg.ProviderConfigs.IpQualityScore.Enabled)
		locator.Show()
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
