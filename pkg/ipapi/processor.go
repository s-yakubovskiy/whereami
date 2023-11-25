package ipapi

import (
	"github.com/s-yakubovskiy/whereami/internal/contracts"
	"github.com/s-yakubovskiy/whereami/internal/whereami"
)

// NOTE: put dto processing logic here
func ConvertIpApiToLocation(ip *IpApiLocation) (*contracts.Location, error) {
	return &contracts.Location{
		IP:          ip.Query,
		Country:     ip.Country,
		CountryCode: ip.CountryCode,
		Region:      ip.Region,
		RegionCode:  ip.RegionName,
		City:        ip.City,
		Timezone:    ip.Timezone,
		Zip:         ip.Zip,
		// Flag:        "",
		Flag:      whereami.CountryCodeToEmoji(ip.CountryCode),
		Isp:       ip.ISP, // Assuming ASN Name represents the ISP
		Org:       ip.Org,
		Latitude:  ip.Lat,
		Longitude: ip.Lon,
		Date:      "", // Set this to current date or as required
		Vpn:       false,
	}, nil
}
