package cmd

import (
	"fmt"
	"log"
	"os"

	migrations "github.com/s-yakubovskiy/whereami"
	"github.com/s-yakubovskiy/whereami/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the WhereAmI application",
	Long:  `This command initializes the WhereAmI application, setting up the SQLite database and applying necessary migrations.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg

		// Initialize the database
		gw, err := migrations.NewGooseWorker(cfg.Database.Path)
		if err != nil {
			log.Fatalf("Failed to create database: %v\n", err)
		}
		if err := gw.InitDB(); err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}

		introduce()
		fmt.Println("Initialization complete.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
