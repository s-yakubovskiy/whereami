package keeper

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/internal/data/db"
	"github.com/s-yakubovskiy/whereami/internal/entity"
	"github.com/s-yakubovskiy/whereami/internal/logging"
)

var _ LocationKeeperRepo = &db.LocationKeeper{}

type UseCase struct {
	cfg *config.AppConfig
	log logging.Logger
	lk  LocationKeeperRepo
}

type LocationData struct {
	Data []entity.Location `json:"data"`
}

type LocationKeeperDumperRepo interface {
	GetAllLocations() ([]entity.Location, error)
	ImportLocations([]entity.Location) error
}

type LocationKeeperVpnRepo interface {
	AddVPNInterface(string) error
	ListVPNInterfaces() ([]string, error)
}

type LocationKeeperRepo interface {
	LocationKeeperVpnRepo
	LocationKeeperDumperRepo
	InitDb(ctx context.Context) error
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

func (uc *UseCase) ExportLocations(ctx context.Context, exportPath string) error {
	uc.log.Debug("UseCase ExportLocations")
	data, err := uc.lk.GetAllLocations()
	if err != nil {
		return err
	}

	// Convert to JSON
	jsonData, err := json.Marshal(LocationData{Data: data})
	if err != nil {
		log.Fatal(err)
	}

	// Write JSON data to file
	err = os.WriteFile(exportPath, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (uc *UseCase) ImportLocations(ctx context.Context, importPath string) error {
	uc.log.Debug("UseCase ImportLocations")

	file, err := os.ReadFile(importPath)
	if err != nil {
		log.Fatal(err)
	}

	var data LocationData
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	if err := uc.lk.ImportLocations(data.Data); err != nil {
		log.Fatal(err)
	}
	return nil
}
