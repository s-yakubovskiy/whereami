package ifconfig

type IfconfigMeMock struct {
	url string
}

func NewPublicIpProviderMock() (*IfconfigMeMock, error) {
	return &IfconfigMeMock{url: "mock"}, nil
}

func (p *IfconfigMeMock) ShowIpProvider() string {
	return p.url
}

func (p *IfconfigMeMock) GetIP() (string, error) {
	return "127.0.0.1", nil
}
