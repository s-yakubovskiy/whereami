// listvpn.go
package main

import (
	"log"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/s-yakubovskiy/whereami/internal/dumper"
	"github.com/spf13/cobra"
)

var importDumpCmd = &cobra.Command{
	Use:   "import",
	Short: "Import all locations from json to sqlite",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg
		importPath := args[0]
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		dumper, err := dumper.NewDumperJSON(dbcli)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		dumper.Import(importPath)
	},
}

func init() {
	dumpCmd.AddCommand(importDumpCmd)
}
