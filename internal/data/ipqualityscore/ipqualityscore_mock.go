package ipqualityscore

import (
	"github.com/s-yakubovskiy/whereami/internal/entity"
	// "github.com/s-yakubovskiy/whereami/internal/contracts"
)

// var _ IpQualityScore = &IpQualityScore{}

// type IpQualityScoreRepo interface {
// 	LookupIpQualityScore(string) (*entity.IpQualityScoreInfo, error)
// }

type IpQualityScoreMock struct {
	url     string
	api_key string
}

func NewIpQualityScoreMock(url, apikey string) (*IpQualityScoreMock, error) {
	return &IpQualityScoreMock{
		url:     url,
		api_key: apikey,
	}, nil
}

func (api *IpQualityScoreMock) LookupIpQualityScore(ip string) (*entity.IpQualityScoreInfo, error) {
	return &entity.IpQualityScoreInfo{
		Success:     true,
		Message:     "hello mock",
		FraudScore:  88,
		CountryCode: "RU",
	}, nil
}
