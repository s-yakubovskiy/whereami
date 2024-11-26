// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/data"
	"github.com/s-yakubovskiy/whereami/internal/data/ifconfig"
	"github.com/s-yakubovskiy/whereami/internal/data/vpn"
	"github.com/s-yakubovskiy/whereami/internal/logging"
	"github.com/s-yakubovskiy/whereami/internal/metrics"
	"github.com/s-yakubovskiy/whereami/internal/server"
	"github.com/s-yakubovskiy/whereami/internal/service"
	"github.com/s-yakubovskiy/whereami/internal/usecase/keeper"
	"github.com/s-yakubovskiy/whereami/internal/usecase/locator"
)

// Injectors from wire.go:

// InitializeShowApp real implementation
func initializeRealShowApp() (*App, func(), error) {
	appConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, nil, err
	}
	logger := logging.ProvideLogger(appConfig)
	ifconfigMe, err := ifconfig.NewPublicIpProvider()
	if err != nil {
		return nil, nil, err
	}
	providerConfigs := config.ProvideProviderConfigs(appConfig)
	ipApi, err := data.ProvideIpApi(providerConfigs)
	if err != nil {
		return nil, nil, err
	}
	ipQualityScore, err := data.ProvideIpQualityScore(providerConfigs)
	if err != nil {
		return nil, nil, err
	}
	metricsMetrics := metrics.NewPrometheusMetrics()
	useCase := locator.NewLocatorUserCase(logger, appConfig, ifconfigMe, ipApi, ipQualityScore, metricsMetrics)
	locationShowService := service.NewLocationShowService(useCase)
	locationKeeper, err := data.ProvideLocationKeeper(appConfig)
	if err != nil {
		return nil, nil, err
	}
	netLinksLister := vpn.NewNetLinkLister()
	keeperUseCase := keeper.NewLocationKeeperUseCase(logger, appConfig, locationKeeper, netLinksLister, metricsMetrics)
	locationKeeperService := service.NewLocationKeeperService(keeperUseCase)
	configServer := config.ProvideServerConfig(appConfig)
	zoshService := service.NewZoshService()
	grpcSrv, err := server.NewGrpcSrv(configServer, logger, zoshService, locationShowService)
	if err != nil {
		return nil, nil, err
	}
	app := NewShowApp(logger, appConfig, locationShowService, locationKeeperService, grpcSrv)
	return app, func() {
	}, nil
}

// InitializeShowApp mock implementation
func initializeMockShowApp() (*App, func(), error) {
	appConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, nil, err
	}
	logger := logging.ProvideLogger(appConfig)
	ifconfigMeMock, err := ifconfig.NewPublicIpProviderMock()
	if err != nil {
		return nil, nil, err
	}
	providerConfigs := config.ProvideProviderConfigs(appConfig)
	ipApiMock, err := data.ProvideIpApiMock(providerConfigs)
	if err != nil {
		return nil, nil, err
	}
	ipQualityScoreMock, err := data.ProvideIpQualityScoreMock(providerConfigs)
	if err != nil {
		return nil, nil, err
	}
	metricsMetrics := metrics.NewPrometheusMetrics()
	useCase := locator.NewLocatorUserCase(logger, appConfig, ifconfigMeMock, ipApiMock, ipQualityScoreMock, metricsMetrics)
	locationShowService := service.NewLocationShowService(useCase)
	locationKeeper, err := data.ProvideLocationKeeper(appConfig)
	if err != nil {
		return nil, nil, err
	}
	netLinksLister := vpn.NewNetLinkLister()
	keeperUseCase := keeper.NewLocationKeeperUseCase(logger, appConfig, locationKeeper, netLinksLister, metricsMetrics)
	locationKeeperService := service.NewLocationKeeperService(keeperUseCase)
	configServer := config.ProvideServerConfig(appConfig)
	zoshService := service.NewZoshService()
	grpcSrv, err := server.NewGrpcSrv(configServer, logger, zoshService, locationShowService)
	if err != nil {
		return nil, nil, err
	}
	app := NewShowApp(logger, appConfig, locationShowService, locationKeeperService, grpcSrv)
	return app, func() {
	}, nil
}

// wire.go:

// InitializeApp sets up the full application with all dependencies.
func InitializeShowApp(isMock bool) (*App, func(), error) {
	if isMock {
		return initializeMockShowApp()
	}
	return initializeRealShowApp()
}
