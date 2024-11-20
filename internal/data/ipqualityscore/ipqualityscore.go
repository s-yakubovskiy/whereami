package ipqualityscore

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/s-yakubovskiy/whereami/internal/entity"
	// "github.com/s-yakubovskiy/whereami/internal/contracts"
)

// var _ IpQualityScore = &IpQualityScore{}

// type IpQualityScoreRepo interface {
// 	LookupIpQualityScore(string) (*entity.IpQualityScoreInfo, error)
// }

type IpQualityScore struct {
	url     string
	api_key string
}

func NewIpQualityScore(url, apikey string) (*IpQualityScore, error) {
	return &IpQualityScore{
		url:     url,
		api_key: apikey,
	}, nil
}

func (api *IpQualityScore) LookupIpQualityScore(ip string) (*entity.IpQualityScoreInfo, error) {
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

	var result *entity.IpQualityScoreInfo
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
