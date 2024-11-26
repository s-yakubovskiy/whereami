package vpn

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

// GetVPN checks if any of the provided VPN interfaces are active on the system.
type NetLinksLister struct{}

func NewNetLinkLister() *NetLinksLister {
	return &NetLinksLister{}
}

func (l *NetLinksLister) GetVPN(vpninterfaces []string) (bool, error) {
	links, err := netlink.LinkList()
	if err != nil {
		return false, fmt.Errorf("error fetching network interfaces: %w", err)
	}

	vpnInterfaceMap := make(map[string]struct{})
	for _, vpnInterface := range vpninterfaces {
		vpnInterfaceMap[vpnInterface] = struct{}{}
	}

	for _, link := range links {
		if _, exists := vpnInterfaceMap[link.Attrs().Name]; exists {
			return true, nil
		}
	}

	return false, nil // No VPN interface found among the system's network interfaces
}
