package ipapi

import "github.com/s-yakubovskiy/whereami/internal/entity"

var _ IpInfoRepo = &IpApi{}

// type IpInfoRepo interface {
// 	LookupIpInfo(string) (*entity.IpInfo, error)
// }

type IpApiMock struct {
	url     string
	api_key string
}

func NewIpApiMock(url, apikey string) (*IpApiMock, error) {
	return &IpApiMock{
		url:     url,
		api_key: apikey,
	}, nil
}

func (l *IpApiMock) LookupIpInfo(ip string) (*entity.Location, error) {
	return &entity.Location{
		Country:     "Russia",
		Isp:         "mock",
		CountryCode: "RU",
		Region:      "Perm",
		City:        "Perm",
		Zip:         "000000",
	}, nil
}
