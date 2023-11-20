package whereami

import (
	"github.com/s-yakubovskiy/whereami/pkg/apiclient"
	"github.com/s-yakubovskiy/whereami/pkg/contracts"
	"github.com/s-yakubovskiy/whereami/pkg/dbclient"
)

// var _ LocatorInterface = &Locator{}

type LocatorInterface interface {
	GetIP() (string, error)
	GetLocation(ip string) (*contracts.Location, error)
	GetVPN() (bool, error)
}

// KeeperInterface defines the interface for database operations
type KeeperInterface interface {
	StoreLocation(location *contracts.Location) error
}

type Locator struct {
	client   LocatorInterface
	dbclient KeeperInterface
}

func NewLocator(api *apiclient.APIClient, dbapi *dbclient.LocationKeeper) *Locator {
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
	location.Output(
		"country",
		"regionname",
		"city",
		"timezone",
		"ip",
	)

	// NOTE: full output avaiable
	// location.Output(
	// 	"Status",
	// 	"Country",
	// 	"CountryCode",
	// 	"Region",
	// 	"RegionName",
	// 	"Zip",
	// 	"City",
	// 	"Lat",
	// 	"Lon",
	// 	"Timezone",
	// 	"Isp",
	// 	"Org",
	// 	"As",
	// 	"IP",
	// 	"Date",
	// )
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
	if err := l.dbclient.StoreLocation(location); err != nil {
		if err.Error() == "The database is already contains this record." {
			warnln(err.Error())
		} else {
			errorln(err.Error())
		}
	}
}
