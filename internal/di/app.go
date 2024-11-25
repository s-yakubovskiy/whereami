package di

import (
	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/logging"
	"github.com/s-yakubovskiy/whereami/internal/service"
)

// AppOption defines a function type that configures App.
type AppOption func(*App) error

type App struct {
	Log            logging.Logger
	Config         *config.AppConfig
	LocatorService *service.LocationShowService
	Keeper         *service.LocationKeeperService
}

// Run is a placeholder for the main application logic.
func (app *App) Run() {
	// Add the main logic of your app here
}

// NewApp creates a new instance of App with all necessary services.
func NewShowApp(
	logging logging.Logger,
	cfg *config.AppConfig,
	locator *service.LocationShowService,
	lk *service.LocationKeeperService,
) *App {
	return &App{
		Config:         cfg,
		Log:            logging,
		LocatorService: locator,
		Keeper:         lk,
	}
}
