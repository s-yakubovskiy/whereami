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
	ipconfig  IPConfigInterface
	primary   IPLocationInterface
	secondary IPLocationInterface
	ipquality IPQualityInterface
}

func NewAPIManager(ip *ipconfig.IPConfig, primary, secondary IPLocationInterface, ipquality IPQualityInterface) *APIManager {
	return &APIManager{
		ipconfig:  ip,
		primary:   primary,
		secondary: secondary,
		ipquality: ipquality,
	}
}

func (l *APIManager) GetIP() (string, error) {
	return l.ipconfig.GetIP()
}

func (l *APIManager) GetLocation(ip string) (*contracts.Location, error) {
	location, err := l.primary.GetLocation(ip)
	if err != nil && l.secondary != nil {
		// NOTE: If the primary client fails, try the secondary client
		return l.secondary.GetLocation(ip)
	}
	return location, err
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

func (l *APIManager) AddIPQuality(location *contracts.Location, ip string) (*contracts.Location, error) {
	return l.ipquality.AddIPQuality(location, ip)
}
