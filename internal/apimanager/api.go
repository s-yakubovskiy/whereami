package apimanager

import (
	"fmt"
	"log"

	"github.com/s-yakubovskiy/whereami/internal/contracts"
	"github.com/s-yakubovskiy/whereami/pkg/ipconfig"

	"github.com/vishvananda/netlink"
)

// IPConfigInterface defines the method for retrieving the current IP address.
type IPConfigInterface interface {
	GetIP() (string, error)
	ShowIpProvider() string
}

// IPLocationInterface defines the method for retrieving the geographical location of an IP address.
type IPLocationInterface interface {
	GetLocation(ip string) (*contracts.Location, error)
}

// APIManager orchestrates operations related to IP configurations, IP location lookups, and IP quality assessments.
type APIManager struct {
	ipconfig  IPConfigInterface
	primary   IPLocationInterface
	secondary IPLocationInterface
	ipquality IPQualityInterface
}

// NewAPIManager creates a new APIManager with specified IP configuration, primary and secondary location services, and an IP quality service.
func NewAPIManager(ip *ipconfig.IPConfig, primary, secondary IPLocationInterface, ipquality IPQualityInterface) *APIManager {
	return &APIManager{
		ipconfig:  ip,
		primary:   primary,
		secondary: secondary,
		ipquality: ipquality,
	}
}

// GetIP proxies the request to the underlying IP configuration service to retrieve the current IP address.
func (l *APIManager) GetIP() (string, error) {
	return l.ipconfig.GetIP()
}

func (l *APIManager) ShowIpProvider() string {
	return l.ipconfig.ShowIpProvider()
}

// GetLocation attempts to find the geographical location of the given IP address, using the primary service and falling back to the secondary if necessary.
func (l *APIManager) GetLocation(ip string) (*contracts.Location, error) {
	location, err := l.primary.GetLocation(ip)
	if err != nil && l.secondary != nil {
		log.Println("Primary client failed. Switching to secondary")
		// If the primary client fails, try the secondary client
		return l.secondary.GetLocation(ip)
	}
	return location, err
}

// GetVPN checks if any of the provided VPN interfaces are active on the system.
func (l *APIManager) GetVPN(vpninterfaces []string) (bool, error) {
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

// AddIPQuality enriches the provided location with quality metrics for the given IP address, using the IP quality service.
func (l *APIManager) AddIPQuality(location *contracts.Location, ip string) (*contracts.Location, error) {
	return l.ipquality.AddIPQuality(location, ip)
}
