package ipapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/s-yakubovskiy/whereami/internal/entity"
)

var _ IpInfoRepo = &IpApi{}

type IpInfoRepo interface {
	LookupIpInfo(string) (*entity.Location, error)
}

type IpApi struct {
	url     string
	api_key string
}

type IpInfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
	Query       string  `json:"query"`
}

func NewIpApi(url, apikey string) (*IpApi, error) {
	return &IpApi{
		url:     url,
		api_key: apikey,
	}, nil
}

func (l *IpApi) LookupIpInfo(ip string) (*entity.Location, error) {
	resp, err := http.Get(l.url + string(ip))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ipInfo *IpInfo
	json.Unmarshal([]byte(jsonData), &ipInfo)
	return ConvertIpApiToLocation(ipInfo)
}

func ConvertIpApiToLocation(ip *IpInfo) (*entity.Location, error) {
	return &entity.Location{
		IP:          ip.Query,
		Country:     ip.Country,
		CountryCode: ip.CountryCode,
		Region:      ip.RegionName,
		RegionCode:  ip.Region,
		City:        ip.City,
		Timezone:    ip.Timezone,
		Zip:         ip.Zip,
		Flag:        CountryCodeToEmoji(ip.CountryCode),
		Isp:         ip.ISP, // Assuming ASN Name represents the ISP
		Latitude:    ip.Lat,
		Longitude:   ip.Lon,
		Date:        "", // Set this to current date or as required
		Vpn:         false,
		Comment:     "Fetched with ipapi provider",
	}, nil
}

func CountryCodeToEmoji(code string) string {
	const offset = 127397 // Offset between uppercase ASCII and regional indicator symbols
	var emoji strings.Builder

	for _, r := range strings.ToUpper(code) {
		if r < 'A' || r > 'Z' {
			return "" // Invalid code
		}
		emoji.WriteRune(rune(r) + offset)
	}

	return emoji.String()
}
