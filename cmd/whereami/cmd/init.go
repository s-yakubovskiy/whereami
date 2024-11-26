package cmd

import (
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
