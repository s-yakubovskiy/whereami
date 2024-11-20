// listvpn.go
package cmd

import (
	"github.com/spf13/cobra"
)

var exportDumpCmd = &cobra.Command{
	Use:   "export",
	Short: "Export all locations to json",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// cfg := config.Cfg
		// exportPath := args[0]
		// dbcli, err := dbclient.NewSQLiteDB(cfg.Database)
		// dumper, err := dumper.NewDumperJSON(dbcli)
		// if err != nil {
		// 	log.Fatalf("Failed to open database: %v", err)
		// }
		// introduce()
		// dumper.Export(exportPath)
	},
}

func init() {
	dumpCmd.AddCommand(exportDumpCmd) // Add to the vpn parent command
}
