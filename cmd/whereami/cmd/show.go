package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var ASYNC_TIMEOUT = 35 * time.Second

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show WhereAmI application",
	Long:  `This command show current public ip address and fetching location information. Print to stdout`,
	Run:   RunShow,
}

// RunShow is the entry point for executing the show command.
// It initializes the application, applies configuration options, and
// starts the location service. It also handles any initialization errors
// and ensures resources are properly cleaned up.
func RunShow(cmd *cobra.Command, args []string) {
	app, ctx, cleanup, err := initializeApp(cmd)
	if err != nil {
		return
	}
	defer cleanup()

	app.LocatorService.ShowLocation(ctx, nil)
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&publicIP, "ip", "i", "", "Specify public IP to lookup info")
}
