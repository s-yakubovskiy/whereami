// vpn.go
package cmd

import (
	pb "github.com/s-yakubovskiy/whereami/api/whrmi/v1"
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Manage dumps JSON <-> sqlite",
	Long:  `Commands to manage dumps including export and import json <-> sqlite3`,
}

var exportDumpCmd = &cobra.Command{
	Use:   "export",
	Short: "Export all locations to json",
	Args:  cobra.ExactArgs(1),
	Run:   dumpExportRun,
}

var importDumpCmd = &cobra.Command{
	Use:   "import",
	Short: "Import all locations from json to sqlite",
	Args:  cobra.ExactArgs(1),
	Run:   dumpImportRun,
}

func dumpExportRun(cmd *cobra.Command, args []string) {
	app, ctx, cleanup, err := initializeApp(cmd)
	if err != nil {
		return
	}
	defer cleanup()

	_, err = app.Keeper.ExportLocations(ctx, &pb.ExportLocationsRequest{Exportpath: args[0]})
	if err != nil {
		app.Log.Fatalf("Failed to retrieve VPN interfaces: %v", err)
	}
}

func dumpImportRun(cmd *cobra.Command, args []string) {
	app, ctx, cleanup, err := initializeApp(cmd)
	if err != nil {
		return
	}
	defer cleanup()

	_, err = app.Keeper.ExportLocations(ctx, &pb.ExportLocationsRequest{Exportpath: args[0]})
	if err != nil {
		app.Log.Fatalf("Failed to retrieve VPN interfaces: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(dumpCmd)
	dumpCmd.AddCommand(exportDumpCmd) // Add to the vpn parent command
	dumpCmd.AddCommand(importDumpCmd)
}
