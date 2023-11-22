package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/apiclient"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/s-yakubovskiy/whereami/internal/dumper"
	"github.com/s-yakubovskiy/whereami/internal/whereami"
)

var fullOutput bool

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show WhereAmI application",
	Long:  `This command Show current public ip address and fetching location information. Print to stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg
		client := apiclient.NewAPIClient()
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		dumper, err := dumper.NewDumperJSON(dbcli)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		locator := whereami.NewLocator(client, dbcli, dumper)
		if fullOutput {
			locator.ShowFull()
		} else {
			locator.Show()
		}
	},
}

func init() {
	showCmd.Flags().BoolVarP(&fullOutput, "full", "f", false, "Display full output")
	rootCmd.AddCommand(showCmd)
}
