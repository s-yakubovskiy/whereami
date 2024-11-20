package contracts

import "github.com/stratoberry/go-gpsd"

// IpProviderInterface defines the method for retrieving the current IP address.
type IpProviderInterface interface {
	GetIP() (string, error)
	ShowIpProvider() string
}

// IPLocationInterface defines the method for retrieving the geographical location of an IP address.
type IPLocationInterface interface {
	GetLocation(ip string) (*Location, error)
}

// IPQualityInterface defines a method for adding quality metrics to a given IP location.
type IPQualityInterface interface {
	AddIPQuality(ip string) (*LocationScores, error)
}

type GPSInterface interface {
	Connect() error
	Close() error
	Fetch() (*gpsd.TPVReport, error)
}

// LKInterface defines methods for interacting with the database regarding location operations.
type LKInterface interface {
	DumperInterface
	StoreLocation(location *Location) error
	ShowLocations(num int) ([]*Location, error)
	AddVPNInterface(interfaceName string) error
	GetVPNInterfaces() ([]string, error)
}

// DumperInterface
type DumperInterface interface {
	GetAllLocations() ([]Location, error)
	ImportLocations([]Location) error
}

// Ensure LocationKeeper implements DBInterface.
// var _ DBInterface = &LocationKeeper{}
