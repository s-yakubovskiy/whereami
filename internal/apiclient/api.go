package apiclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/s-yakubovskiy/whereami/internal/contracts"

	"github.com/vishvananda/netlink"
)

const (
	IP_CONFIG_ADDR   = "http://ifconfig.me"
	IP_LOCATION_ADDR = "http://ip-api.com/json/"
)

type APIClient struct{}

func NewAPIClient() *APIClient {
	return &APIClient{}
}

func (l *APIClient) GetIP() (string, error) {
	// // Fetching public IP address
	resp, err := http.Get(IP_CONFIG_ADDR)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}

func (l *APIClient) GetLocation(ip string) (*contracts.Location, error) {
	resp, err := http.Get(IP_LOCATION_ADDR + string(ip))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *contracts.Location
	json.Unmarshal([]byte(jsonData), &result)

	return result, nil
}

func (l *APIClient) GetVPN(vpninterfaces []string) (bool, error) {
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
