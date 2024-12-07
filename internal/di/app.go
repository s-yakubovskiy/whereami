package di

import (
	"context"

	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/server"
	"github.com/s-yakubovskiy/whereami/internal/service"
	"github.com/s-yakubovskiy/whereami/pkg/shudralogs"
)

// AppOption defines a function type that configures App.
type AppOption func(*App) error

type App struct {
	Log            shudralogs.Logger
	Config         *config.AppConfig
	LocatorService *service.LocationShowService
	Keeper         *service.LocationKeeperService
	Gs             *server.GrpcGatewaySrv
}

// Run is a placeholder for the main application logic.
func (app *App) Run() {
	// Add the main logic of your app here
}

func (a *App) NewContext() context.Context {
	ctx := context.Background()
	return ctx
}

// NewApp creates a new instance of App with all necessary services.
func NewShowApp(
	logging shudralogs.Logger,
	cfg *config.AppConfig,
	locator *service.LocationShowService,
	lk *service.LocationKeeperService,
	gs *server.GrpcGatewaySrv,
) *App {
	return &App{
		Config:         cfg,
		Log:            logging,
		LocatorService: locator,
		Keeper:         lk,
		Gs:             gs,
	}
}
