package ipconfig

import (
	"io"
	"net/http"
)

const (
	IP_CONFIG_ADDR = "http://ifconfig.me"
)

type IPConfig struct{}

func NewIPConfig() (*IPConfig, error) {
	return &IPConfig{}, nil
}

func (l *IPConfig) GetIP() (string, error) {
	// // Fetching public IP address
	resp, err := http.Get(IP_CONFIG_ADDR)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}
