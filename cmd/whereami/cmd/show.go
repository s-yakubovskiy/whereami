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
	"github.com/spf13/cobra"
)

var (
	fullShow    bool
	locationApi string
	publicIpApi string
	publicIP    string
	gpsEnabled  bool
	gpsProvider string
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show WhereAmI application",
	Long:  `This command Show current public ip address and fetching location information. Print to stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := &config.Cfg
		if locationApi != "" {
			cfg.MainProvider = locationApi
		}
		if publicIpApi != "" {
			cfg.ProviderConfigs.PublicIpProvider = publicIpApi
		}
		if publicIP != "" {
			cfg.PublicIP = publicIP
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
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database)
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

		locator := whereami.NewLocator(client, dbcli, dumper, gps, cfg)
		introduce()
		if fullShow {
			locator.ShowFull()
		} else {
			locator.Show()
		}
	},
}

func init() {
	showCmd.Flags().BoolVarP(&fullShow, "full", "f", false, "Display full output")
	showCmd.Flags().StringVarP(&locationApi, "location-api", "l", "", "Select ip location provider: [ipapi, ipdata]")
	showCmd.Flags().StringVarP(&publicIpApi, "public-ip-api", "p", "", "Select public ip api provider: [ifconfig.me, ipinfo.io, icanhazip.com]")
	showCmd.Flags().StringVarP(&publicIP, "ip", "i", "", "Specify public IP to lookup info")
	showCmd.Flags().BoolVarP(&gpsEnabled, "gps", "", false, "Add experimental GPS integration [gpsd service should be up & running]")
	rootCmd.AddCommand(showCmd)
}
