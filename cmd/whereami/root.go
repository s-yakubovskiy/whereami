package main

import (
	"fmt"
	"os"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/pkg/apiclient"
	"github.com/s-yakubovskiy/whereami/pkg/dbclient"
	"github.com/s-yakubovskiy/whereami/pkg/whereami"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whereami",
	Short: "WhereAmI is an application to find your geolocation based on your IP",
	Long:  `WhereAmI is a CLI application that allows users to find their geolocation based on their public IP address.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg
		client := apiclient.NewAPIClient()
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open database: %v\n", err)
			os.Exit(1)
		}
		locator := whereami.NewLocator(client, dbcli)
		locator.Show()
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
