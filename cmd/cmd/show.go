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

var (
	fullShow     bool
	providerShow string
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show WhereAmI application",
	Long:  `This command Show current public ip address and fetching location information. Print to stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg
		if providerShow != "" {
			cfg.MainProvider = providerShow
		}

		ipconfig, err := ipconfig.NewIPConfig()
		primary, secondary, ipquality, err := getLocationClient(cfg.MainProvider)

		client := apimanager.NewAPIManager(ipconfig, primary, secondary, ipquality)
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		dumper, err := dumper.NewDumperJSON(dbcli)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		locator := whereami.NewLocator(client, dbcli, dumper, cfg.ProviderConfigs.IpQualityScore.Enabled)
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
	showCmd.Flags().StringVarP(&providerShow, "provider", "p", "", "Select ip location provider: [ipapi, ipdata]")
	rootCmd.AddCommand(showCmd)
}
