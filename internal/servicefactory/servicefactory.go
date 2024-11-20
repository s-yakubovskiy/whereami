package servicefactory

import (
	"fmt"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/apimanager"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
	"github.com/s-yakubovskiy/whereami/pkg/gpsdfetcher"
	// "github.com/s-yakubovskiy/whereami/pkg/ipapi"
)

// ServiceFactory interface declares methods for creating service clients
type ServiceFactory interface {
	CreateLocationService(cfg config.ProviderConfig) (contracts.IPLocationInterface, error)
	CreateQualityService(cfg config.ProviderConfig) (contracts.IPQualityInterface, error)
	CreateIpProviderService(cfg config.ProviderConfig) (contracts.IpProviderInterface, error)
}

// DefaultServiceFactory struct implements the ServiceFactory interface
type DefaultServiceFactory struct{}

func NewDefaultServiceFactory() *DefaultServiceFactory {
	return &DefaultServiceFactory{}
}

// CreateLocationService creates a location service client based on the configuration
func (f *DefaultServiceFactory) CreateLocationService(cfg config.ProviderConfig) (contracts.IPLocationInterface, error) {
	switch cfg.Name {
	case "ipapi":
		// return ipapi.NewIpApi(cfg.URL, cfg.APIKey)
		return apimanager.NewIpApiClient(cfg)
	case "ipdata":
		return apimanager.NewIpDataClient(cfg)
	default:
		return nil, fmt.Errorf("unknown location service provider: %s", cfg.Name)
	}
}

// CreateQualityService creates an IP quality service client if enabled
func (f *DefaultServiceFactory) CreateQualityService(cfg config.ProviderConfig) (contracts.IPQualityInterface, error) {
	if cfg.Enabled {
		return apimanager.NewIpQualityScoreClient(cfg)
		// return ipqualityscore.NewIpQualityScore(cfg.URL, cfg.APIKey)
	}
	return nil, fmt.Errorf("IP quality service not enabled")
}

// CreateIpProviderService creates an IP provider service client
func (f *DefaultServiceFactory) CreateIpProviderService(cfg config.ProviderConfig) (contracts.IpProviderInterface, error) {
	switch cfg.Name {
	case "ifconfig":
		return apimanager.NewIpProviderClient(cfg)
	default:
		return nil, fmt.Errorf("unknown ip provider service: %s", cfg.Name)
	}
}

func (f *DefaultServiceFactory) ProvideGPSFetcher(cfg *config.GPSConfig) contracts.GPSInterface {
	if cfg.Enabled {
		switch cfg.Provider {
		case "adb":
			return gpsdfetcher.NewGPSAdbFetcher()
		case "file":
			return gpsdfetcher.NewGPSDFileFetcher(cfg.Timeout)
		default:
			return gpsdfetcher.NewGPSDFetcher(cfg.Timeout)
		}
	}
	return nil // or some default GPS fetcher
}
