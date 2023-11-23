package ipapi

import (
	ipdata "github.com/ipdata/go"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

// NOTE: put dto processing logic here
func ConvertIPToLocation(ip ipdata.IP) *contracts.Location {
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
	}
}
