package cmd

import (
	"log"

	"github.com/robfig/cron/v3"
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

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run WhereAmI as a background service",
	Long:  `This command starts the WhereAmI service as a daemon that performs tasks based on the crontab configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		startDaemon()
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}

func startDaemon() {
	cfg := config.Cfg
	factory := &servicefactory.DefaultServiceFactory{}

	c := cron.New()
	for _, task := range cfg.CrontabTasks {
		taskCopy := task // Create a copy of the task for the current iteration
		_, err := c.AddFunc(taskCopy.Schedule, func() {
			ipconfig, err := ipconfig.NewIPConfig(cfg.ProviderConfigs.PublicIpProvider)
			if err != nil {
				log.Printf("Failed to create IP configuration: %v", err)
				return
			}

			ipapi, err := factory.CreateLocationService(cfg.ProviderConfigs.IpApi)
			if err != nil {
				log.Printf("Failed to create primary location service: %v", err)
				return
			}
			ipdata, err := factory.CreateLocationService(cfg.ProviderConfigs.IpData)
			if err != nil {
				log.Printf("Failed to create secondary location service: %v", err)
				return
			}
			ipquality, err := factory.CreateQualityService(cfg.ProviderConfigs.IpQualityScore)
			if err != nil {
				log.Printf("Failed to create IP quality service: %v", err)
				return
			}

			client := apimanager.NewAPIManager(ipconfig, ipapi, ipdata, ipquality)
			dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
			if err != nil {
				log.Printf("Failed to open database: %v", err)
				return
			}
			dumper, err := dumper.NewDumperJSON(dbcli)
			if err != nil {
				log.Printf("Failed to create dumper: %v", err)
				return
			}
			var gps gpsdfetcher.GPSInterface
			if cfg.GPSConfig.Enabled || gpsEnabled {
				cfg.GPSConfig.Enabled = true
				gps = gpsdfetcher.NewGPSDFetcher(cfg.GPSConfig.Timeout)
			}

			lCfg := whereami.NewConfig(cfg.ProviderConfigs.IpQualityScore.Enabled, ipLookup, gpsEnabled)
			locator := whereami.NewLocator(client, dbcli, dumper, gps, lCfg)
			locator.Store()
		})
		if err != nil {
			log.Printf("Error scheduling task: %v", err)
		}
	}

	c.Start()

	// Block this goroutine, as c.Start() runs in its own goroutine.
	select {}
}
