package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func serverRun(cmd *cobra.Command, args []string) {
	app, ctx, cleanup, err := initializeApp(cmd)
	if err != nil {
		return
	}
	defer cleanup()

	// Handle shutdown signals
	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		app.Log.Info("Shutdown signal received!")
		cancel()
	}()

	// Start serving
	if err := app.Gs.Serve(ctx); err != nil {
		app.Log.Errorf("failed to serve: %v", err)
	}

	app.Log.Info("Server shutdown gracefully")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server the WhereAmI application",
	Long:  `Just servering.`,
	Run:   serverRun,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
