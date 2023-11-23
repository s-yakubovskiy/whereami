package apimanager

import (
	"fmt"

	"github.com/s-yakubovskiy/whereami/internal/contracts"
	"github.com/s-yakubovskiy/whereami/pkg/ipconfig"

	"github.com/vishvananda/netlink"
)

type IPConfigInterface interface {
	GetIP() (string, error)
}
type IPLocationInterface interface {
	GetLocation(ip string) (*contracts.Location, error)
}

type APIManager struct {
	ipconfig IPConfigInterface
	api      IPLocationInterface
}

func NewAPIManager(ip *ipconfig.IPConfig, api IPLocationInterface) *APIManager {
	return &APIManager{
		ipconfig: ip,
		api:      api,
	}
}

func (l *APIManager) GetIP() (string, error) {
	return l.ipconfig.GetIP()
}

func (l *APIManager) GetLocation(ip string) (*contracts.Location, error) {
	return l.api.GetLocation(ip)
}

func (l *APIManager) GetVPN(vpninterfaces []string) (bool, error) {
	// Fetch all network interfaces
	links, err := netlink.LinkList()
	if err != nil {
		return false, fmt.Errorf("error fetching network interfaces: %w", err)
	}

	// Create a map for efficient lookup
	vpnInterfaceMap := make(map[string]struct{})
	for _, vpnInterface := range vpninterfaces {
		vpnInterfaceMap[vpnInterface] = struct{}{}
	}

	// Check if any of the system's interfaces are in vpninterfaces
	for _, link := range links {
		if _, exists := vpnInterfaceMap[link.Attrs().Name]; exists {
			return true, nil
		}
	}

	return false, nil // No VPN interface found in the system's network interfaces
}
