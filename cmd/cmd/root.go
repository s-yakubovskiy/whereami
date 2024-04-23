package cmd

import (
	"fmt"
	"log"
	"os"

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
			if gpsProvider == "adb" {
				gps = gpsdfetcher.NewGPSAdbFetcher()
			} else if gpsProvider == "file" {
				gps = gpsdfetcher.NewGPSDFileFetcher(cfg.GPSConfig.Timeout)
			} else {
				gps = gpsdfetcher.NewGPSDFetcher(cfg.GPSConfig.Timeout)
			}
		}

		lCfg := whereami.NewConfig(cfg.ProviderConfigs.IpQualityScore.Enabled, ipLookup, cfg.GPSConfig.Enabled)
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
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Display application version")
	rootCmd.Flags().BoolVarP(&gpsEnabled, "gps", "", false, "Add experimental GPS integration [gpsd service should be up & running]")
	rootCmd.Flags().StringVarP(&locationApi, "location-api", "l", "", "Select IP location provider: [ipapi, ipdata]")
	rootCmd.Flags().StringVarP(&publicIpApi, "public-ip-api", "p", "", "Select public IP API provider: [ifconfig.me, ipinfo.io, icanhazip.com]")
	rootCmd.Flags().StringVarP(&gpsProvider, "gps-provider", "g", "", "Select GPS provider: [adb, file, gpsd (default)]")
	rootCmd.Flags().StringVarP(&ipLookup, "ip", "", "", "Specify public IP to lookup info")
}
