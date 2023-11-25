package ipapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

type IpApiClient struct {
	url     string
	api_key string
}

type IpApiLocation struct {
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

func NewIpApiClient(providerConfig config.ProviderConfig) (*IpApiClient, error) {
	return &IpApiClient{
		url:     providerConfig.URL,
		api_key: providerConfig.APIKey,
	}, nil
}

func (l *IpApiClient) GetLocation(ip string) (*contracts.Location, error) {
	resp, err := http.Get(l.url + string(ip))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *IpApiLocation
	json.Unmarshal([]byte(jsonData), &result)

	return ConvertIpApiToLocation(result)
}
