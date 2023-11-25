package ipqualityscore

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IpQualityScore struct {
	url     string
	api_key string
}

type IpQualityScoreLocation struct {
	Success        bool    `json:"success"`
	Message        string  `json:"message"`
	FraudScore     int     `json:"fraud_score"`
	CountryCode    string  `json:"country_code"`
	Region         string  `json:"region"`
	City           string  `json:"city"`
	ISP            string  `json:"ISP"`
	ASN            int     `json:"ASN"`
	Organization   string  `json:"organization"`
	IsCrawler      bool    `json:"is_crawler"`
	Timezone       string  `json:"timezone"`
	Mobile         bool    `json:"mobile"`
	Host           string  `json:"host"`
	Proxy          bool    `json:"proxy"`
	VPN            bool    `json:"vpn"`
	Tor            bool    `json:"tor"`
	ActiveVPN      bool    `json:"active_vpn"`
	ActiveTor      bool    `json:"active_tor"`
	RecentAbuse    bool    `json:"recent_abuse"`
	BotStatus      bool    `json:"bot_status"`
	ConnectionType string  `json:"connection_type"`
	AbuseVelocity  string  `json:"abuse_velocity"`
	ZipCode        string  `json:"zip_code"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	RequestId      string  `json:"request_id"`
}

func NewIpQualityScore(url, apikey string) (*IpQualityScore, error) {
	return &IpQualityScore{
		url:     url,
		api_key: apikey,
	}, nil
}

func (api *IpQualityScore) Lookup(ip string) (*IpQualityScoreLocation, error) {
	requestURL := fmt.Sprintf("%s/%s/%s?strictness=2&fast=0", api.url, api.api_key, ip)
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *IpQualityScoreLocation
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
