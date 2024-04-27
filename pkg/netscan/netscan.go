// Package netscan provides a simple network scanner for open TCP ports.
package netscan

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// Scanner defines the interface for network scanning.
type Scanner interface {
	ScanPort(ip string, port int) bool
	ScanSubnet(subnet string, port int) []string
	GetSubnets() ([]string, error)
}

// TCPScanner implements the Scanner interface for TCP port scanning.
type TCPScanner struct {
	Timeout time.Duration
}

// NewTCPScanner creates a new TCPScanner with the specified timeout.
func NewTCPScanner(timeout time.Duration) *TCPScanner {
	return &TCPScanner{
		Timeout: timeout,
	}
}

// ScanPort checks if a TCP port is open on a given IP.
func (t *TCPScanner) ScanPort(ip string, port int) bool {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, t.Timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// ScanSubnet scans all IPs in a CIDR subnet and returns a list of IPs with the specified port open.
func (t *TCPScanner) ScanSubnet(subnet string, port int) []string {
	ip, ipnet, err := net.ParseCIDR(subnet)
	if err != nil {
		fmt.Println("Parse error:", err)
		return nil
	}

	var openHosts []string
	var wg sync.WaitGroup
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			if t.ScanPort(ip, port) {
				openHosts = append(openHosts, ip)
			}
		}(ip.String())
	}
	wg.Wait()
	return openHosts
}

// GetSubnets returns a slice of subnets available on active WLAN network interfaces.
func (t *TCPScanner) GetSubnets() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var subnets []string
	for _, iface := range interfaces {
		// Check for WLAN interface names commonly starting with 'wl' or 'wlan'
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 || (!strings.HasPrefix(iface.Name, "wl") && !strings.HasPrefix(iface.Name, "wlan")) {
			continue // skip down, loopback, or non-WLAN interfaces
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
				subnets = append(subnets, ipnet.String())
			}
		}
	}
	return subnets, nil
}

// incIP increments the IP address by 1.
func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
