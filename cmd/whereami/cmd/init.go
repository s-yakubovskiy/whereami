package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func initRun(cmd *cobra.Command, args []string) {
	app, ctx, cleanup, err := initializeApp(cmd)
	if err != nil {
		return
	}
	defer cleanup()

	app.Keeper.Init(ctx, nil)
	app.Log.Info("Initialization complete")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the WhereAmI application",
	Long:  `This command initializes the WhereAmI application, setting up the SQLite database and applying necessary migrations.`,
	Run:   initRun,
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
