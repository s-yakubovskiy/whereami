// vpn.go
package cmd

import (
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Manage dumps JSON <-> sqlite",
	Long:  `Commands to manage dumps including export and import json <-> sqlite3`,
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
