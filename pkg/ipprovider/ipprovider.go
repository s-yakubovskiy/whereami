package ipprovider

import (
	"encoding/json"
	"io"
	"net/http"
)

type IPProvider struct {
	url string
}

type IPInfo struct {
	IPAddr     string `json:"ip_addr"`
	RemoteHost string `json:"remote_host"`
	UserAgent  string `json:"user_agent"`
	Port       string `json:"port"`
	Method     string `json:"method"`
	MIME       string `json:"mime"`
	Via        string `json:"via"`
	Forwarded  string `json:"forwarded"`
}

func NewIPProvider(url string) (*IPProvider, error) {
	return &IPProvider{url: url}, nil
}

func (p *IPProvider) ShowIpProvider() string {
	return p.url
}

func (p *IPProvider) GetIP() (string, error) {
	resp, err := http.Get(p.url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var info IPInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		return "", err
	}

	return info.IPAddr, nil
}
