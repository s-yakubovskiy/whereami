package data

import (
	"github.com/google/wire"
	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/data/ifconfig"
	"github.com/s-yakubovskiy/whereami/internal/data/ipapi"
	"github.com/s-yakubovskiy/whereami/internal/data/ipqualityscore"
)

func ProvideIpApi(config *config.ProviderConfigs) (*ipapi.IpApi, error) {
	c, _ := config.GetConfig("ipapi")
	return ipapi.NewIpApi(c.URL, c.APIKey)
}

func ProvideIpApiMock(config *config.ProviderConfigs) (*ipapi.IpApiMock, error) {
	c, _ := config.GetConfig("ipapi")
	return ipapi.NewIpApiMock(c.URL, c.APIKey)
}

func ProvideIpQualityScore(config *config.ProviderConfigs) (*ipqualityscore.IpQualityScore, error) {
	c, _ := config.GetConfig("ipqualityscore")
	return ipqualityscore.NewIpQualityScore(c.URL, c.APIKey)
}

func ProvideIpQualityScoreMock(config *config.ProviderConfigs) (*ipqualityscore.IpQualityScoreMock, error) {
	c, _ := config.GetConfig("ipqualityscore")
	return ipqualityscore.NewIpQualityScoreMock(c.URL, c.APIKey)
}

var ProviderSet = wire.NewSet(
	ifconfig.NewPublicIpProvider,
	ifconfig.NewPublicIpProviderMock,
	ProvideIpApi,
	ProvideIpApiMock,
	ProvideIpQualityScore,
	ProvideIpQualityScoreMock)
