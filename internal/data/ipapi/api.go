package ipapi

import (
	"encoding/json"
	"io"
	"net/http"
)

var _ IpInfoRepo = &IpApi{}

type IpInfoRepo interface {
	LookupIpInfo(string) (*IpInfo, error)
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

func (l *IpApi) LookupIpInfo(ip string) (*IpInfo, error) {
	resp, err := http.Get(l.url + string(ip))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *IpInfo
	json.Unmarshal([]byte(jsonData), &result)
	return result, err
}
