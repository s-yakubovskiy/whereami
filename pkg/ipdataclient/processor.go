package ipdataclient

import (
	"fmt"

	ipdata "github.com/ipdata/go"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

// NOTE: put dto processing logic here
func ConvertIPToLocation(ip ipdata.IP) (*contracts.Location, error) {
	// Check for nil pointers in fields that are pointers in the source struct
	if ip.TimeZone == nil {
		return nil, fmt.Errorf("missing required fields in IP data")
	}
	return &contracts.Location{
		IP:          ip.IP,
		Country:     ip.CountryName,
		CountryCode: ip.CountryCode,
		Region:      ip.Region,
		RegionCode:  ip.RegionCode,
		City:        ip.City,
		Timezone:    ip.TimeZone.Name,
		Zip:         ip.Postal,
		Flag:        ip.EmojiFlag,
		// EmojiFlag:   ip.EmojiFlag,
		Isp:       ip.Flag, // Assuming ASN Name represents the ISP
		Org:       ip.Organization,
		Latitude:  ip.Latitude,
		Longitude: ip.Longitude,
		Date:      "", // Set this to current date or as required
		Vpn:       false,
	}, nil
}
