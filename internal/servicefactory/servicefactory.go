package servicefactory

import (
	"fmt"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/apimanager"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

// ServiceFactory interface declares methods for creating service clients
type ServiceFactory interface {
	CreateLocationService(cfg config.ProviderConfig) (contracts.IPLocationInterface, error)
	CreateQualityService(cfg config.ProviderConfig) (contracts.IPQualityInterface, error)
}

// DefaultServiceFactory struct implements the ServiceFactory interface
type DefaultServiceFactory struct{}

// CreateLocationService creates a location service client based on the configuration
func (f *DefaultServiceFactory) CreateLocationService(cfg config.ProviderConfig) (contracts.IPLocationInterface, error) {
	switch cfg.Name {
	case "ipapi":
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
	}
	return nil, fmt.Errorf("IP quality service not enabled")
}
