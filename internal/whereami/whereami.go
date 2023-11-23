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
}

// KeeperInterface defines the interface for database operations
type KeeperInterface interface {
	StoreLocation(location *contracts.Location) error
	AddVPNInterface(interfaceName string) error
	GetVPNInterfaces() ([]string, error)
}

type Locator struct {
	client   LocatorInterface
	dbclient KeeperInterface
	dumper   dumper.DumperJSON
}

func NewLocator(api *apimanager.APIManager, dbapi *dbclient.LocationKeeper, dumper *dumper.DumperJSON) *Locator {
	return &Locator{
		client:   api,
		dbclient: dbapi,
	}
}

func (l *Locator) Show() {
	// Fetching data from IP API
	ip, err := l.client.GetIP()
	if err != nil {
		errorln(err.Error())
	}
	location, err := l.client.GetLocation(ip)
	if err != nil {
		errorln(err.Error())
	}
	if location != nil && ip != "" {
		vpninterfaces, err := l.dbclient.GetVPNInterfaces()
		if err != nil {
			warnln(err.Error())
		}

		vpn, err := l.client.GetVPN(vpninterfaces)
		if err != nil {
			warnln(err.Error())
		}

		if vpn {
			location.Vpn = true
		}

		// output to stding colorized
		location.Output(
			"country",
			"regionname",
			"city",
			"timezone",
			"ip",
			"vpn",
		)
	}
}

func (l *Locator) ShowFull() {
	// Fetching data from IP API
	ip, err := l.client.GetIP()
	if err != nil {
		errorln(err.Error())
	}
	location, err := l.client.GetLocation(ip)
	if err != nil {
		errorln(err.Error())
	}
	if location != nil && ip != "" {
		vpninterfaces, err := l.dbclient.GetVPNInterfaces()
		if err != nil {
			warnln(err.Error())
		}

		vpn, err := l.client.GetVPN(vpninterfaces)
		if err != nil {
			warnln(err.Error())
		}

		if vpn {
			location.Vpn = true
		}

		// output to stding colorized
		location.Output(
			"status",
			"country",
			"countrycode",
			"region",
			"regionname",
			"zip",
			"city",
			"lat",
			"lon",
			"timezone",
			"isp",
			"org",
			"as",
			"ip",
			"vpn",
		)
	}
}

func (l *Locator) Store() {
	ip, err := l.client.GetIP()
	if err != nil {
		errorln(err.Error())
	}
	location, err := l.client.GetLocation(ip)
	if err != nil {
		errorln(err.Error())
	}

	if location != nil && ip != "" {
		vpninterfaces, err := l.dbclient.GetVPNInterfaces()
		if err != nil {
			warnln(err.Error())
		}

		vpn, err := l.client.GetVPN(vpninterfaces)
		if err != nil {
			warnln(err.Error())
		}
		if vpn {
			location.Vpn = true
		}
		if err := l.dbclient.StoreLocation(location); err != nil {
			if err.Error() == "The database is already contains this record." {
				warnln(err.Error())
			} else {
				errorln(err.Error())
			}
		}
	}
}
