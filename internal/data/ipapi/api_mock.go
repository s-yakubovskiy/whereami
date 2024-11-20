package ipapi

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

func (l *IpApiMock) LookupIpInfo(ip string) (*IpInfo, error) {
	return &IpInfo{
		Org:         "mock",
		AS:          "mock",
		Status:      "mock",
		Country:     "Russia",
		CountryCode: "RU",
		Region:      "Perm",
		RegionName:  "Perm",
		City:        "Perm",
		Zip:         "000000",
	}, nil
}
