// listvpn.go
package main

import (
	"log"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/pkg/dbclient"
	"github.com/s-yakubovskiy/whereami/pkg/dumper"
	"github.com/spf13/cobra"
)

var exportDumpCmd = &cobra.Command{
	Use:   "export",
	Short: "Export all locations to json",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg
		exportPath := args[0]
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		dumper, err := dumper.NewDumperJSON(dbcli)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		dumper.Export(exportPath)
	},
}

func init() {
	dumpCmd.AddCommand(exportDumpCmd) // Add to the vpn parent command
}
