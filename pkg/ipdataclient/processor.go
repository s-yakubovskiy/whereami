package ipdataclient

import (
	"fmt"

	ipdata "github.com/ipdata/go"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

// NOTE: put dto processing logic here
func (l *IPDataClient) ConvertIPToLocation(ip ipdata.IP) (*contracts.Location, error) {
	// Check for nil pointers in fields that are pointers in the source struct
	if ip.TimeZone == nil {
		return nil, fmt.Errorf("missing required fields in IP data")
	}
	return &contracts.Location{
		Status:      "success",
		Country:     ip.CountryName,
		CountryCode: ip.CountryCode,
		Region:      ip.RegionCode,
		RegionName:  ip.Region,
		City:        ip.City,
		Zip:         ip.Postal,
		Lat:         ip.Latitude,
		Lon:         ip.Longitude,
		Timezone:    ip.TimeZone.Name,
		Isp:         ip.ASN.Name, // Assuming ASN Name represents the ISP
		Org:         ip.Organization,
		As:          ip.ASN.ASN, // Using ASN as the 'as' field
		IP:          ip.IP,
		Date:        "", // Set this to current date or as required
		Vpn:         false,
	}, nil
}
