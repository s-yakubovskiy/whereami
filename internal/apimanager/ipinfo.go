package apimanager

import (
	"log"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

type FallbackLocationService struct {
	Primary   contracts.IPLocationInterface
	Secondary contracts.IPLocationInterface
}

func NewFallbackLocationService(primary, secondary contracts.IPLocationInterface) *FallbackLocationService {
	return &FallbackLocationService{
		Primary:   primary,
		Secondary: secondary,
	}
}

func (f *FallbackLocationService) GetLocation(ip string) (*contracts.Location, error) {
	location, err := f.Primary.GetLocation(ip)
	if err != nil && f.Secondary != nil {
		log.Println("Primary location service failed, switching to secondary")
		return f.Secondary.GetLocation(ip)
	}
	return location, err
}

// IpOption is an HTTP server option.
type IpOption func(*IpInfoClient)

// IpData with ipinfoclient ipdata client.
func IpDataOption(cfg config.ProviderConfig) IpOption {
	return func(s *IpInfoClient) {
		s.ipdata, _ = NewIpDataClient(cfg)
	}
}

// IpApi with ipinfoclient ipapi client.
func IpApiOption(cfg config.ProviderConfig) IpOption {
	return func(s *IpInfoClient) {
		s.ipapi, _ = NewIpApiClient(cfg)
	}
}

// IpInfolient wraps the ipapi client and implements the IPLocationInterface for getting location information.
type IpInfoClient struct {
	ipapi  *IpApiClient
	ipdata *IpDataClient
}

// NewApplyIpOption creates a new instance of IpApiClient using the provided configuration.
func NewIpOptionClient(cfgs []config.ProviderConfig) (*IpInfoClient, error) {
	var opts []IpOption

	for _, cfg := range cfgs {
		if cfg.Name == "ipdata" {
			opts = append(opts, IpDataOption(cfg))
		}
		if cfg.Name == "ipapi" {
			opts = append(opts, IpApiOption(cfg))
		}
	}

	client, err := NewApplyIpOption(opts...)
	return client, err
}

// NewApplyIpOption creates a new instance of IpApiClient using the provided configuration.
func NewApplyIpOption(opts ...IpOption) (*IpInfoClient, error) {
	srv := &IpInfoClient{}
	for _, o := range opts {
		o(srv)
	}
	return srv, nil
}

func (f *IpInfoClient) GetLocation(ip string) (*contracts.Location, error) {
	location, err := f.ipapi.GetLocation(ip)
	if err != nil && f.ipdata != nil {
		log.Println("Primary location service failed, switching to secondary")
		return f.ipdata.GetLocation(ip)
	}
	return location, err
}
