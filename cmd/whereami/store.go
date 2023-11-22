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

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store WhereAmI application",
	Long:  `This command store location information to datbase (sqlite)`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg
		client := apiclient.NewAPIClient()
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		dumper, err := dumper.NewDumperJSON(dbcli)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		locator := whereami.NewLocator(client, dbcli, dumper)
		locator.Store()
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
