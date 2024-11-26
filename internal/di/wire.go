//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/data"
	"github.com/s-yakubovskiy/whereami/internal/data/db"
	"github.com/s-yakubovskiy/whereami/internal/data/ifconfig"
	"github.com/s-yakubovskiy/whereami/internal/data/ipapi"
	"github.com/s-yakubovskiy/whereami/internal/data/ipqualityscore"
	"github.com/s-yakubovskiy/whereami/internal/data/vpn"
	"github.com/s-yakubovskiy/whereami/internal/logging"
	"github.com/s-yakubovskiy/whereami/internal/server"
	"github.com/s-yakubovskiy/whereami/internal/service"
	"github.com/s-yakubovskiy/whereami/internal/usecase"
	"github.com/s-yakubovskiy/whereami/internal/usecase/keeper"
	"github.com/s-yakubovskiy/whereami/internal/usecase/locator"
)

// InitializeApp sets up the full application with all dependencies.
func InitializeShowApp(isMock bool) (*App, func(), error) {
	if isMock {
		return initializeMockShowApp()
	}
	return initializeRealShowApp()
}

// InitializeShowApp real implementation
func initializeRealShowApp() (*App, func(), error) {
	wire.Build(
		config.ProviderSet,
		logging.ProvideLogger,
		service.ProviderSet,
		usecase.ProviderSet,
		data.ProviderSet,
		server.ProviderSet,

		wire.Bind(new(keeper.NetLinksRepo), new(*vpn.NetLinksLister)),
		wire.Bind(new(locator.PublicIpRepo), new(*ifconfig.IfconfigMe)),
		wire.Bind(new(locator.IpInfoRepo), new(*ipapi.IpApi)),
		wire.Bind(new(keeper.LocationKeeperRepo), new(*db.LocationKeeper)),
		wire.Bind(new(locator.IpQualityScoreRepo), new(*ipqualityscore.IpQualityScore)),
		NewShowApp,
	)
	return &App{}, nil, nil
}

// InitializeShowApp mock implementation
func initializeMockShowApp() (*App, func(), error) {
	wire.Build(
		config.ProviderSet,
		logging.ProvideLogger,
		service.ProviderSet,
		usecase.ProviderSet,
		data.ProviderSet,
		server.ProviderSet,

		wire.Bind(new(keeper.NetLinksRepo), new(*vpn.NetLinksLister)),
		wire.Bind(new(locator.PublicIpRepo), new(*ifconfig.IfconfigMeMock)),
		wire.Bind(new(locator.IpInfoRepo), new(*ipapi.IpApiMock)),
		wire.Bind(new(keeper.LocationKeeperRepo), new(*db.LocationKeeper)),

		wire.Bind(new(locator.IpQualityScoreRepo), new(*ipqualityscore.IpQualityScoreMock)),
		NewShowApp,
	)
	return &App{}, nil, nil
}
