package ipconfig

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type IPConfig struct {
	url string
}

func NewIPConfig(url string) (*IPConfig, error) {
	return &IPConfig{url: fmt.Sprintf("http://%s", url)}, nil
}

func (l *IPConfig) ShowIpProvider() string {
	return l.url
}

func (l *IPConfig) GetIP() (string, error) {
	// // Fetching public IP address
	resp, err := http.Get(l.url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.Join(strings.Fields(string(ip)), " "), nil
}
