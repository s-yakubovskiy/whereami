package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/apimanager"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/s-yakubovskiy/whereami/internal/dumper"
	"github.com/s-yakubovskiy/whereami/internal/whereami"
	"github.com/s-yakubovskiy/whereami/pkg/gpsdfetcher"
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
		introduce()
		if showVersion {
			fmt.Println("\nBuild Info:")
			fmt.Println("  Version:", Version)
			fmt.Println("  Commit:", Commit)
			os.Exit(0)
		}
		if locationApi != "" {
			cfg.MainProvider = locationApi
		}
		if publicIpApi != "" {
			cfg.ProviderConfigs.PublicIpProvider = publicIpApi
		}

		ipconfig, err := ipconfig.NewIPConfig(cfg.ProviderConfigs.PublicIpProvider)
		primary, secondary, ipquality, err := getLocationClient(cfg.MainProvider)

		client := apimanager.NewAPIManager(ipconfig, primary, secondary, ipquality)
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		dumper, err := dumper.NewDumperJSON(dbcli)
		gps := gpsdfetcher.NewGPSDFetcher(cfg.GPSConfig.Timeout)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		lCfg := whereami.NewConfig(cfg.ProviderConfigs.IpQualityScore.Enabled, ipLookup, gpsEnabled)
		locator := whereami.NewLocator(client, dbcli, dumper, gps, lCfg)
		if fullShow {
			locator.ShowFull()
		} else {
			locator.Show()
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&fullShow, "full", "f", false, "Display full output")
	rootCmd.Flags().StringVarP(&locationApi, "location-api", "l", "", "Select ip location provider: [ipapi, ipdata]")
	rootCmd.Flags().StringVarP(&publicIpApi, "public-ip-api", "p", "", "Select public ip api provider: [ifconfig.me, ipinfo.io, icanhazip.com]")
	rootCmd.Flags().StringVarP(&ipLookup, "ip", "", "", "Specify public IP to lookup info")
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Display application version")
	//
}
