package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/apimanager"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/s-yakubovskiy/whereami/internal/dumper"
	"github.com/s-yakubovskiy/whereami/internal/whereami"
	"github.com/s-yakubovskiy/whereami/pkg/gpsdfetcher"
	"github.com/s-yakubovskiy/whereami/pkg/ipconfig"
)

var (
	fullShow    bool
	locationApi string
	publicIpApi string
	ipLookup    string
	gpsEnabled  bool
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show WhereAmI application",
	Long:  `This command Show current public ip address and fetching location information. Print to stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg
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
		if cfg.GPSConfig.Enabled {
			gpsEnabled = true
		}
		lCfg := whereami.NewConfig(cfg.ProviderConfigs.IpQualityScore.Enabled, ipLookup, gpsEnabled)
		locator := whereami.NewLocator(client, dbcli, dumper, gps, lCfg)
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
	showCmd.Flags().StringVarP(&ipLookup, "ip", "i", "", "Specify public IP to lookup info")
	showCmd.Flags().BoolVarP(&gpsEnabled, "gps", "", false, "Add experimental GPS intergration [gpsd service should up & running]")
	rootCmd.AddCommand(showCmd)
}
