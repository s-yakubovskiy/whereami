package whereami

import (
	"github.com/s-yakubovskiy/whereami/internal/apimanager"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/s-yakubovskiy/whereami/internal/dumper"
)

// var _ LocatorInterface = &Locator{}

type LocatorInterface interface {
	GetLocation(ip string) (*contracts.Location, error)
	GetVPN([]string) (bool, error)
	GetIP() (string, error)
	ShowIpProvider() string
	AddIPQuality(*contracts.Location, string) (*contracts.Location, error)
}

// KeeperInterface defines the interface for database operations
type KeeperInterface interface {
	StoreLocation(location *contracts.Location) error
	AddVPNInterface(interfaceName string) error
	GetVPNInterfaces() ([]string, error)
}

type Locator struct {
	client    LocatorInterface
	dbclient  KeeperInterface
	dumper    *dumper.DumperJSON
	ipquality bool
	iplookup  string
}

type Config struct {
	IpQuality bool
	IP        string
}

func NewConfig(ipquality bool, ip string) *Config {
	return &Config{
		IpQuality: ipquality,
		IP:        ip,
	}
}

func NewLocator(api *apimanager.APIManager, dbapi *dbclient.LocationKeeper, dumper *dumper.DumperJSON, cfg *Config) *Locator {
	return &Locator{
		client:    api,
		dbclient:  dbapi,
		dumper:    dumper,
		ipquality: cfg.IpQuality,
		iplookup:  cfg.IP,
	}
}
