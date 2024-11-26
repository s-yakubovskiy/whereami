package cmd

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/robfig/cron/v3"
	"github.com/s-yakubovskiy/whereami/api/whrmi/v1"
	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/di"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run WhereAmI as a background service",
	Long:  `This command starts the WhereAmI service as a daemon that performs tasks based on the crontab configuration.`,
	Run:   daemonRun,
}

var (
	cronScheduler *cron.Cron
	initOnce      sync.Once
	app           *di.App
)

func init() {
	rootCmd.AddCommand(daemonCmd)
}

func initializeAppOnce(cmd *cobra.Command) error {
	var err error
	initOnce.Do(func() {
		app, _, _, err = initializeApp(cmd)
	})
	go app.Gs.ServeSync()
	return err
}

// Main function to run daemon
func daemonRun(cmd *cobra.Command, args []string) {
	cfg := config.Cfg

	showIntro = false
	err := initializeAppOnce(cmd)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	cronScheduler = cron.New()
	for _, task := range cfg.CrontabTasks {
		taskCopy := task // Copy the task for the current iteration

		_, err := cronScheduler.AddFunc(taskCopy.Schedule, func() {
			performTask(taskCopy)
		})
		if err != nil {
			log.Printf("Error scheduling task: %v", err)
		}
	}

	cronScheduler.Start()
	signalHandler()
}

// Function to fetch location and store
func performTask(task config.CrontabTask) { // assuming task is of some TaskType
	ctx := app.NewContext() // Hypothetical context getter
	data, err := app.LocatorService.GetLocation(ctx, nil)
	if err != nil {
		app.Log.Printf("Failed to retrieve location: %v", err)
		return
	}

	_, err = app.Keeper.StoreLocation(ctx, &whrmi.StoreLocationRequest{Location: data.Location})
	if err != nil {
		app.Log.Printf("Failed to save location: %v", err)
	}
	ctx.Done()
}

// Handle system signals for graceful shutdown
func signalHandler() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	<-signalChannel
	log.Println("Received shutdown signal")
	cronScheduler.Stop() // Stop cron scheduler

	// Additional cleanup if necessary
	os.Exit(0)
}
