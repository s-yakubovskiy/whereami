// listvpn.go
package main

import (
	"fmt"
	"log"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/spf13/cobra"
)

var listVPNCmd = &cobra.Command{
	Use:   "list",
	Short: "List all network interfaces marked as VPN",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Cfg
		dbcli, err := dbclient.NewSQLiteDB(cfg.Database.Path)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		interfaces, err := dbcli.GetVPNInterfaces()
		if err != nil {
			log.Fatalf("Failed to retrieve VPN interfaces: %v", err)
		}

		if len(interfaces) == 0 {
			fmt.Println("No VPN interfaces found.")
			return
		}

		fmt.Println("VPN Interfaces:")
		for _, iface := range interfaces {
			fmt.Println(" -", iface)
		}
	},
}

func init() {
	vpnCmd.AddCommand(listVPNCmd) // Add to the vpn parent command
}
