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
	nl  NetLinksRepo
}

type NetLinksRepo interface {
	GetVPN([]string) (bool, error)
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
	StoreLocation(*entity.Location) error
}

func NewLocationKeeperUseCase(log logging.Logger, cfg *config.AppConfig, lk LocationKeeperRepo, nl NetLinksRepo) *UseCase {
	return &UseCase{
		cfg: cfg,
		log: log,
		lk:  lk,
		nl:  nl,
	}
}

func (uc *UseCase) GetVPN(vpninterfaces []string) (bool, error) {
	return uc.nl.GetVPN(vpninterfaces)
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

func (uc *UseCase) StoreLocation(ctx context.Context, location *entity.Location) error {
	uc.log.Debug("UseCase StoreLocation")
	vpninterfaces, err := uc.lk.ListVPNInterfaces()
	if err != nil {
		uc.log.Error(err.Error())
	}

	vpn, err := uc.nl.GetVPN(vpninterfaces)
	if err != nil {
		uc.log.Error(err.Error())
	}
	if vpn {
		location.Vpn = true
	}
	// TODO: update it later
	// if l.cfg.IpQuality {
	// 	l.client.AddIPQuality(location, ip)
	// }
	if err := uc.lk.StoreLocation(location); err != nil {
		if err.Error() == "The database already contains this record." {
			uc.log.Warn(err.Error())
		} else {
			uc.log.Error(err.Error())
		}
	}
	return nil
}
