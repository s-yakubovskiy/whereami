// vpn.go
package main

import (
	"github.com/spf13/cobra"
)

var vpnCmd = &cobra.Command{
	Use:   "vpn",
	Short: "Manage VPN interfaces",
	Long:  `Commands to manage VPN interfaces, including listing and marking them.`,
}

func init() {
	rootCmd.AddCommand(vpnCmd)
}
