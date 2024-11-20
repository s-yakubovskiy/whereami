package whereami

import (
	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/apimanager"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
	"github.com/s-yakubovskiy/whereami/internal/dbclient"
	"github.com/s-yakubovskiy/whereami/internal/dumper"
)

var _ LocatorInterface = &apimanager.APIManager{}

type LocatorInterface interface {
	GetLocation(ip string) (*contracts.Location, error)
	GetVPN([]string) (bool, error)
	GetIP() (string, error)
	ShowIpProvider() string
	AddIPQuality(string) (*contracts.LocationScores, error)
}

// KeeperInterface defines the interface for database operations
type KeeperInterface interface {
	StoreLocation(location *contracts.Location) error
	AddVPNInterface(interfaceName string) error
	GetVPNInterfaces() ([]string, error)
	ShowLocations(num int) ([]*contracts.Location, error)
}

type Locator struct {
	client   LocatorInterface
	dbclient KeeperInterface
	dumper   *dumper.DumperJSON
	gps      contracts.GPSInterface
	cfg      *config.AppConfig
}

type Config struct {
	IpQuality bool
	IP        string
	GPS       bool
}

func NewLocator(api *apimanager.APIManager, dbapi *dbclient.LocationKeeper, dumper *dumper.DumperJSON, gps contracts.GPSInterface, cfg *config.AppConfig) *Locator {
	return &Locator{
		client:   api,
		dbclient: dbapi,
		dumper:   dumper,
		gps:      gps,
		cfg:      cfg,
	}
}
