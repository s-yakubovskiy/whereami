package cmd

import (
	"fmt"
	"log"

	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the WhereAmI application",
	Long:  `This command initializes the WhereAmI application, setting up the SQLite database and applying necessary migrations.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbcli, err := dbclient.NewSQLiteDB("~/work/common/whereami_locations.sqlite")
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		// Initialize the database
		if err := dbcli.InitDB(); err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}

		fmt.Println("Initialization complete.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
