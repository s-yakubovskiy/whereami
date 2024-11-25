package cmd

import (
	"github.com/s-yakubovskiy/whereami/api/whrmi/v1"
	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store WhereAmI application",
	Long:  `This command stores location information in the database (sqlite).`,
	Run:   storeRun,
}

func storeRun(cmd *cobra.Command, args []string) {
	app, ctx, cleanup, err := initializeApp(cmd)
	if err != nil {
		return
	}
	defer cleanup()

	data, err := app.LocatorService.GetLocation(ctx, nil)
	if err != nil {
		app.Log.Fatalf("Failed to retrieve location: %v", err)
	}
	_, err = app.Keeper.StoreLocation(ctx, &whrmi.StoreLocationRequest{Location: data.Location})
	if err != nil {
		app.Log.Fatalf("Failed to save location: %v", err)
	}
}

func init() {
	storeCmd.Flags().StringVarP(&locationApi, "provider", "p", "", "Select IP location provider: [ipapi, ipdata]")
	rootCmd.AddCommand(storeCmd)
}
