package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
)

var markVPNCmd = &cobra.Command{
	Use:   "mark [interface]",
	Short: "Mark a network interface as VPN",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg
		interfaceName := args[0]
		introduce()

		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		err = dbcli.AddVPNInterface(interfaceName)
		if err != nil {
			log.Fatalf("Failed to mark interface as VPN: %v", err)
		}

		fmt.Printf("Interface %s marked as VPN successfully\n", interfaceName)
	},
}

func init() {
	vpnCmd.AddCommand(markVPNCmd)
}
