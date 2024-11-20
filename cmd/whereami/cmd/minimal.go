package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/di"
	"github.com/spf13/cobra"
)

var ASYNC_TIMEOUT = 35 * time.Second

var minimalCmd = &cobra.Command{
	Use:   "minimal",
	Short: "minimal WhereAmI application",
	Long:  `This command minimal current public ip address and fetching location information. Print to stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		isMockEnv := os.Getenv("USE_MOCK") == "true"
		app, cleanup, err := di.InitializeShowApp(isMockEnv)
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error initializing application: %v\n", err)
			return
		}

		config.ApplyOptions(
			config.WithPublicIP(publicIP),
		)
		RunMinimal(app)
		defer cleanup()
	},
}

// Separate the command execution logic into a testable function
func RunMinimal(app *di.App) {
	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), ASYNC_TIMEOUT)
	defer cancel()
	introduce()

	app.LocatorService.ShowLocation(ctx, nil)
}

func init() {
	rootCmd.AddCommand(minimalCmd)
	minimalCmd.Flags().StringVarP(&publicIP, "ip", "i", "", "Specify public IP to lookup info")
}
