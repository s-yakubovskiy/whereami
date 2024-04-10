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

var NumHistory int32

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "History of WhereAmI locations",
	Long:  `This command output History of wheremai locations. Print to stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		introduce()
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
		locator.History(NumHistory)
	},
}

func init() {
	historyCmd.Flags().Int32VarP(&NumHistory, "number", "n", 10, "Specify number of last history locations to output")
	rootCmd.AddCommand(historyCmd)
}
