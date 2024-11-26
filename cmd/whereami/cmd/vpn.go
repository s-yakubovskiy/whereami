// vpn.go
package cmd

import (
	"fmt"

	pb "github.com/s-yakubovskiy/whereami/api/whrmi/v1"
	"github.com/spf13/cobra"
)

var vpnCmd = &cobra.Command{
	Use:   "vpn",
	Short: "Manage VPN interfaces",
	Long:  `Commands to manage VPN interfaces, including listing and marking them.`,
}

var listVPNCmd = &cobra.Command{
	Use:   "list",
	Short: "List all network interfaces marked as VPN",
	Run:   vpnListRun,
}

var markVPNCmd = &cobra.Command{
	Use:   "mark [interface]",
	Short: "Mark a network interface as VPN",
	Args:  cobra.ExactArgs(1),
	Run:   vpnMarkRun,
}

func vpnMarkRun(cmd *cobra.Command, args []string) {
	app, ctx, cleanup, err := initializeApp(cmd)
	if err != nil {
		return
	}
	defer cleanup()

	_, err = app.Keeper.AddVpnInterface(ctx, &pb.AddVpnInterfaceRequest{Vpninterface: args[0]})
	if err != nil {
		app.Log.Fatalf("Failed to retrieve VPN interfaces: %v", err)
	}
}

func vpnListRun(cmd *cobra.Command, args []string) {
	app, ctx, cleanup, err := initializeApp(cmd)
	if err != nil {
		return
	}
	defer cleanup()

	ifaces, err := app.Keeper.ListVpnInterfaces(ctx, nil)
	if err != nil {
		app.Log.Fatalf("Failed to retrieve VPN interfaces: %v", err)
	}

	if len(ifaces.Vpninterfaces) == 0 {
		fmt.Println("No VPN interfaces found.")
		return
	}

	fmt.Println("VPN Interfaces:")
	for _, iface := range ifaces.Vpninterfaces {
		app.Log.Printf(" - %s", iface)
	}
}

func init() {
	rootCmd.AddCommand(vpnCmd)
	vpnCmd.AddCommand(listVPNCmd) // Add to the vpn parent command
}
