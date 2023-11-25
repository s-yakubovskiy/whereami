package apimanager

import (
	"fmt"
	"strconv"

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

func ConvertIpQualityScoreToLocation(ip *ipqualityscore.IpQualityScoreLocation) (*contracts.Location, error) {
	// Check for nil pointers in fields that are pointers in the source struct
	return &contracts.Location{
		CountryCode: ip.CountryCode,
		Region:      ip.Region,
		City:        ip.City,
		Timezone:    ip.Timezone,
		Zip:         ip.ZipCode,
		Flag:        common.CountryCodeToEmoji(ip.CountryCode),
		Isp:         ip.ISP, // Assuming ASN Name represents the ISP
		Asn:         strconv.Itoa(ip.ASN),
		Latitude:    ip.Latitude,
		Longitude:   ip.Longitude,
		Date:        "", // Set this to current date or as required
		Vpn:         false,
		// Scores: ip.IP,
		Comment: "Fetched with ipqualityscore provider",
	}, nil
}

func EnrichLocationWithQualityScore(location *contracts.Location, ip *ipqualityscore.IpQualityScoreLocation) (*contracts.Location, error) {
	location.Scores.FraudScore = ip.FraudScore
	location.Scores.Host = ip.Host
	location.Scores.IsCrawler = ip.IsCrawler
	location.Scores.BotStatus = ip.BotStatus
	location.Scores.Tor = ip.Tor
	location.Scores.Proxy = ip.Proxy
	location.Scores.VPN = ip.VPN
	location.Comment = location.Comment + ". Updated with ipqualityscore provider"

	return location, nil
}
