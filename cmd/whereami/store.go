package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/s-yakubovskiy/whereami/pkg/apiclient"
	"github.com/s-yakubovskiy/whereami/pkg/dbclient"
	"github.com/s-yakubovskiy/whereami/pkg/whereami"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store WhereAmI application",
	Long:  `This command store location information to datbase (sqlite)`,
	Run: func(cmd *cobra.Command, args []string) {
		client := apiclient.NewAPIClient()
		dbcli, err := dbclient.NewSQLiteDB("~/work/common/whereami_locations.sqlite")
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		locator := whereami.NewLocator(client, dbcli)
		locator.Store()
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
