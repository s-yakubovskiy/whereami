package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/apimanager"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/s-yakubovskiy/whereami/internal/dumper"
	"github.com/s-yakubovskiy/whereami/internal/whereami"
	"github.com/s-yakubovskiy/whereami/pkg/ipconfig"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store WhereAmI application",
	Long:  `This command store location information to datbase (sqlite)`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg
		if locationApi != "" {
			cfg.MainProvider = locationApi
		}
		ipconfig, err := ipconfig.NewIPConfig(cfg.ProviderConfigs.PublicIpProvider)
		primary, secondary, ipquality, err := getLocationClient(cfg.MainProvider)

		client := apimanager.NewAPIManager(ipconfig, primary, secondary, ipquality)
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		dumper, err := dumper.NewDumperJSON(dbcli)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		lCfg := whereami.NewConfig(cfg.ProviderConfigs.IpQualityScore.Enabled, ipLookup)
		locator := whereami.NewLocator(client, dbcli, dumper, lCfg)
		locator.Store()
	},
}

func init() {
	storeCmd.Flags().StringVarP(&locationApi, "provider", "p", "", "Select ip location provider: [ipapi, ipdata]")
	rootCmd.AddCommand(storeCmd)
}
