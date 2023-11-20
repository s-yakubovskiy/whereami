package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/s-yakubovskiy/whereami/pkg/apiclient"
	"github.com/s-yakubovskiy/whereami/pkg/dbclient"
	"github.com/s-yakubovskiy/whereami/pkg/whereami"
)

var getCmd = &cobra.Command{
	Use:   "show",
	Short: "Show WhereAmI application",
	Long:  `This command Show current public ip address and fetching location information. Print to stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		client := apiclient.NewAPIClient()
		dbcli, err := dbclient.NewSQLiteDB("~/work/common/whereami_locations.sqlite")
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		locator := whereami.NewLocator(client, dbcli)
		locator.Show()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
