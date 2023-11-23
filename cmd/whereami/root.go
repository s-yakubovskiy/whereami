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
	"github.com/s-yakubovskiy/whereami/pkg/ipapi"
	"github.com/s-yakubovskiy/whereami/pkg/ipconfig"
	"github.com/s-yakubovskiy/whereami/pkg/ipdataclient"
	"github.com/spf13/cobra"
)

func getLocationClient(provider string) (apimanager.IPLocationInterface, error) {
	switch provider {
	case "ipapi":
		c := config.Cfg.ProviderConfigs.IpApi
		return ipapi.NewIpApiClient(c)
	case "ipdata":
		// NOTE: change to correct config
		c := config.Cfg.ProviderConfigs.IpData
		return ipdataclient.NewIPDataClient(c)
	default:
		log.Fatalf("Unknown provider set: %+v\n", provider)
		return nil, fmt.Errorf("unknown provider: %s", provider)
	}
}

var rootCmd = &cobra.Command{
	Use:   "whereami",
	Short: "WhereAmI is an application to find your geolocation based on your IP",
	Long:  `WhereAmI is a CLI application that allows users to find their geolocation based on their public IP address.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg

		ipconfig, err := ipconfig.NewIPConfig()
		locationApi, err := getLocationClient(cfg.MainProvider)

		client := apimanager.NewAPIManager(ipconfig, locationApi)
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		dumper, err := dumper.NewDumperJSON(dbcli)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open database: %v\n", err)
			os.Exit(1)
		}
		locator := whereami.NewLocator(client, dbcli, dumper)
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
