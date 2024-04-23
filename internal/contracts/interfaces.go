package contracts

// IPConfigInterface defines the method for retrieving the current IP address.
type IPConfigInterface interface {
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
