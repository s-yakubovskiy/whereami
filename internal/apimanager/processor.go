package apimanager

import (
	"fmt"

	ipdata "github.com/ipdata/go"
	"github.com/s-yakubovskiy/whereami/internal/common"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
	"github.com/s-yakubovskiy/whereami/pkg/ipapi"
	"github.com/s-yakubovskiy/whereami/pkg/ipqualityscore"
)

func ConvertIpApiToLocation(ip *ipapi.IpApiLocation) (*contracts.Location, error) {
	return &contracts.Location{
		IP:          ip.Query,
		Country:     ip.Country,
		CountryCode: ip.CountryCode,
		Region:      ip.RegionName,
		RegionCode:  ip.Region,
		City:        ip.City,
		Timezone:    ip.Timezone,
		Zip:         ip.Zip,
		Flag:        common.CountryCodeToEmoji(ip.CountryCode),
		Isp:         ip.ISP, // Assuming ASN Name represents the ISP
		Latitude:    ip.Lat,
		Longitude:   ip.Lon,
		Date:        "", // Set this to current date or as required
		Vpn:         false,
		Comment:     "Fetched with ipapi provider",
	}, nil
}

func ConvertIpDataToLocation(ip ipdata.IP) (*contracts.Location, error) {
	// Check for nil pointers in fields that are pointers in the source struct
	if ip.TimeZone == nil {
		return nil, fmt.Errorf("missing required fields in IP data")
	}
	return &contracts.Location{
		IP: ip.IP,
		// Country: ip.CountryName,
		Country:     ip.CountryName,
		CountryCode: ip.CountryCode,
		Region:      ip.Region,
		RegionCode:  ip.RegionCode,
		City:        ip.City,
		Timezone:    ip.TimeZone.Name,
		Zip:         ip.Postal,
		Flag:        ip.EmojiFlag,
		// EmojiFlag:   ip.EmojiFlag,
		Isp:       ip.ASN.Name, // Assuming ASN Name represents the ISP
		Asn:       ip.ASN.ASN,
		Latitude:  ip.Latitude,
		Longitude: ip.Longitude,
		Date:      "", // Set this to current date or as required
		Vpn:       false,
		// Scores: ip.IP,
		Comment: "Fetched with ipdata provider",
	}, nil
}

func EnrichLocationWithQualityScore(location *contracts.LocationScores, ip *ipqualityscore.IpQualityScoreLocation) (*contracts.LocationScores, error) {
	location.FraudScore = ip.FraudScore
	location.Host = ip.Host
	location.IsCrawler = ip.IsCrawler
	location.BotStatus = ip.BotStatus
	location.Tor = ip.Tor
	location.Proxy = ip.Proxy
	location.VPN = ip.VPN
	return location, nil
}
