package cmd

import (
	"context"
	"fmt"
	"os"

	// "github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/di"
	"github.com/spf13/cobra"
)

// version overrides
var (
	Version     string
	Commit      string
	showVersion bool
)

// public flags for whrmi
var (
	fullShow    bool
	locationApi string
	publicIpApi string
	publicIP    string
	gpsEnabled  bool
	gpsProvider string
)

var rootCmd = &cobra.Command{
	Use:   "whereami",
	Short: "WhereAmI is an application to find your geolocation based on your IP",
	Long:  `WhereAmI is a CLI application that allows users to find their geolocation based on their public IP address.`,
	Run: func(cmd *cobra.Command, args []string) {
		// cfg := &config.Cfg
		introduce()
		if showVersion {
			fmt.Println("\nBuild Info:")
			fmt.Println("  Version:", Version)
			fmt.Println("  Commit:", Commit)
			os.Exit(0)
		}
	},
}

func initializeApp(cmd *cobra.Command) (*di.App, context.Context, context.CancelFunc, error) {
	isMockEnv := os.Getenv("USE_MOCK") == "true"
	app, cleanup, err := di.InitializeShowApp(isMockEnv)
	if err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "Error initializing application: %v\n", err)
		return nil, nil, nil, err
	}

	config.ApplyOptions(
		config.WithPublicIP(publicIP),
	)
	// app.Log.PrettyPrint(app.Config)

	ctx, cancel := context.WithTimeout(context.Background(), ASYNC_TIMEOUT)
	introduce()
	return app, ctx, func() { cleanup(); cancel() }, nil
}

func init() {
	rootCmd.Flags().BoolVarP(&fullShow, "full", "f", false, "Display full output")
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Display application version")
	rootCmd.Flags().BoolVarP(&gpsEnabled, "gps", "", false, "Add experimental GPS integration [gpsd service should be up & running]")
	rootCmd.Flags().StringVarP(&locationApi, "location-api", "l", "", "Select IP location provider: [ipapi, ipdata]")
	rootCmd.Flags().StringVarP(&publicIpApi, "public-ip-api", "p", "", "Select public IP API provider: [ifconfig.me, ipinfo.io, icanhazip.com]")
	rootCmd.Flags().StringVarP(&gpsProvider, "gps-provider", "g", "", "Select GPS provider: [adb, file, gpsd (default)]")
	rootCmd.Flags().StringVarP(&publicIP, "ip", "", "", "Specify public IP to lookup info")
}
