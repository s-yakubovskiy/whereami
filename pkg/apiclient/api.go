package apiclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/s-yakubovskiy/whereami/pkg/contracts"

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

func (l *APIClient) GetVPN() (bool, error) {
	// Fetch all routes from all tables
	routes, err := netlink.RouteList(nil, netlink.FAMILY_ALL)
	if err != nil {
		return false, fmt.Errorf("error fetching routes: %w", err)
	}

	for _, route := range routes {
		fmt.Printf("%+v\n", route.Protocol)
		if route.Dst == nil { // Check for a nil Dst (default route)
			link, err := netlink.LinkByIndex(route.LinkIndex)
			if err != nil {
				return false, fmt.Errorf("error fetching interface for route: %w", err)
			}

			// Check if the interface matches common VPN patterns
			if matched, _ := regexp.MatchString(`^(tun|wg)\d`, link.Attrs().Name); matched {
				if route.Src != nil && route.Protocol == 4 {
					return true, nil // VPN with static preferred source detected
				}
			}
		}
	}

	return false, nil // No VPN Detected
}
