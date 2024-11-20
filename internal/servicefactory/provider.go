package servicefactory

import (
	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

// This function wraps the factory method to be used in Wire setup
func ProvideGPSFetcher(factory *DefaultServiceFactory, cfgGps *config.GPSConfig) contracts.GPSInterface {
	return factory.ProvideGPSFetcher(cfgGps)
}
