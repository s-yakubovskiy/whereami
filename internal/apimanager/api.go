package apimanager

import (
	"fmt"

	"github.com/s-yakubovskiy/whereami/internal/contracts"
	"github.com/vishvananda/netlink"
)

type APIManager struct {
	ipprovider contracts.IpProviderInterface
	location   contracts.IPLocationInterface // Now a single interface
	ipquality  contracts.IPQualityInterface
}

func NewAPIManager(ip contracts.IpProviderInterface, location contracts.IPLocationInterface, ipquality contracts.IPQualityInterface) *APIManager {
	return &APIManager{
		ipprovider: ip,
		location:   location,
		ipquality:  ipquality,
	}
}

func (a *APIManager) GetIP() (string, error) {
	return a.ipprovider.GetIP()
}

func (a *APIManager) ShowIpProvider() string {
	return a.ipprovider.ShowIpProvider()
}

func (a *APIManager) GetLocation(ip string) (*contracts.Location, error) {
	return a.location.GetLocation(ip)
}

func (a *APIManager) GetVPN(vpninterfaces []string) (bool, error) {
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
	return false, nil
}

func (a *APIManager) AddIPQuality(ip string) (*contracts.LocationScores, error) {
	return a.ipquality.AddIPQuality(ip)
}
