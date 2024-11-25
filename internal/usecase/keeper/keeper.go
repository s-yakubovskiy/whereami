package keeper

import (
	"context"

	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/data/db"

	// "github.com/s-yakubovskiy/whereami/internal/contracts"

	"github.com/s-yakubovskiy/whereami/internal/logging"
)

var _ LocationKeeperRepo = &db.LocationKeeper{}

type UseCase struct {
	cfg *config.AppConfig
	log logging.Logger
	lk  LocationKeeperRepo
}

type LocationKeeperRepo interface {
	InitDb(ctx context.Context) error
	AddVPNInterface(string) error
	ListVPNInterfaces() ([]string, error)
}

func NewLocationKeeperUseCase(log logging.Logger, cfg *config.AppConfig, lk LocationKeeperRepo) *UseCase {
	return &UseCase{
		cfg: cfg,
		log: log,
		lk:  lk,
	}
}

func (uc *UseCase) InitDb(ctx context.Context) error {
	uc.log.Debug("UseCase DbService")
	if err := uc.lk.InitDb(ctx); err != nil {
		uc.log.Errorf("Error initializing db: %v", err)
		return err
	}
	return nil
}

func (uc *UseCase) AddVPNInterface(ctx context.Context, vpnInterface string) error {
	uc.log.Debug("UseCase AddVPNInterface")
	if err := uc.lk.AddVPNInterface(vpnInterface); err != nil {
		uc.log.Errorf("Error initializing db: %v", err)
		return err
	}
	return nil
}

func (uc *UseCase) ListVPNInterfaces() ([]string, error) {
	uc.log.Debug("UseCase ListVPNInterface")
	interfaces, err := uc.lk.ListVPNInterfaces()
	if err != nil {
		uc.log.Errorf("Error initializing db: %v", err)
		return nil, err
	}
	return interfaces, nil
}
