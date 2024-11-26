package ifconfig

import (
	"encoding/json"
	"io"
	"net/http"
)

type PublicIpProvider interface {
	ShowIpProvider() string
	GetIP() (string, error)
}

type IfconfigMe struct {
	url string
}

const IPCONFIG_URL = "http://ifconfig.me/all.json"

type IfconfigMeResp struct {
	IPAddr     string `json:"ip_addr"`
	RemoteHost string `json:"remote_host"`
	UserAgent  string `json:"user_agent"`
	Port       string `json:"port"`
	Method     string `json:"method"`
	MIME       string `json:"mime"`
	Via        string `json:"via"`
	Forwarded  string `json:"forwarded"`
}

func NewPublicIpProvider() (*IfconfigMe, error) {
	return &IfconfigMe{url: IPCONFIG_URL}, nil
}

func (p *IfconfigMe) ShowIpProvider() string {
	return p.url
}

func (p *IfconfigMe) GetIP() (string, error) {
	resp, err := http.Get(p.url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var info IfconfigMeResp
	err = json.Unmarshal(body, &info)
	if err != nil {
		return "", err
	}

	return info.IPAddr, nil
}
