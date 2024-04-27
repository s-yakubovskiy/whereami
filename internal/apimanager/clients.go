package apimanager

import (
	ipdata "github.com/ipdata/go"
	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
	"github.com/s-yakubovskiy/whereami/pkg/ipapi"
	"github.com/s-yakubovskiy/whereami/pkg/ipprovider"
	"github.com/s-yakubovskiy/whereami/pkg/ipqualityscore"
)

var _ contracts.IPQualityInterface = &IpQualityScoreClient{}

// IpApiClient wraps the ipapi client and implements the IPLocationInterface for getting location information.
type IpApiClient struct {
	client *ipapi.IpApi
}

// NewIpApiClient creates a new instance of IpApiClient using the provided configuration.
func NewIpApiClient(cfg config.ProviderConfig) (*IpApiClient, error) {
	client, err := ipapi.NewIpApi(cfg.URL, cfg.APIKey)
	return &IpApiClient{
		client: client,
	}, err
}

// GetLocation uses the ipapi client to lookup the location of the given IP and convert it to a Location struct.
func (l *IpApiClient) GetLocation(ip string) (*contracts.Location, error) {
	data, err := l.client.Lookup(ip)
	if err != nil {
		return nil, err
	}
	return ConvertIpApiToLocation(data)
}

// IpDataClient wraps the ipdata client and implements the IPLocationInterface for getting location information.
type IpDataClient struct {
	client *ipdata.Client
}

// NewIpDataClient creates a new instance of IpDataClient using the provided configuration.
func NewIpDataClient(cfg config.ProviderConfig) (*IpDataClient, error) {
	client, err := ipdata.NewClient(cfg.APIKey)
	return &IpDataClient{
		client: &client,
	}, err
}

// GetLocation uses the ipdata client to lookup the location of the given IP and convert it to a Location struct.
func (l *IpDataClient) GetLocation(ip string) (*contracts.Location, error) {
	data, err := l.client.Lookup(ip)
	if err != nil {
		return nil, err
	}
	return ConvertIpDataToLocation(data)
}

// thin wrapper for ipquality
// IpQualityScoreClient wraps the ipqualityscore client and implements the IPQualityInterface for adding quality metrics.
type IpQualityScoreClient struct {
	client *ipqualityscore.IpQualityScore
}

// NewIpQualityScoreClient creates a new instance of IpQualityScoreClient using the provided configuration.
func NewIpQualityScoreClient(cfg config.ProviderConfig) (*IpQualityScoreClient, error) {
	client, err := ipqualityscore.NewIpQualityScore(cfg.URL, cfg.APIKey)
	return &IpQualityScoreClient{
		client: client,
	}, err
}

// AddIPQuality uses the ipqualityscore client to lookup the quality score of the given IP and enriches the Location struct with this data.
func (l *IpQualityScoreClient) AddIPQuality(ip string) (*contracts.LocationScores, error) {
	location := contracts.NewLocationScores()

	qualityScore, err := l.client.Lookup(ip)
	if err != nil {
		return nil, err
	}

	return EnrichLocationWithQualityScore(location, qualityScore)
}

// thin wrapper around ipprovider
type IpProviderClient struct {
	client *ipprovider.IPProvider
}

func NewIpProviderClient(cfg config.ProviderConfig) (*IpProviderClient, error) {
	client, err := ipprovider.NewIPProvider(cfg.URL)
	return &IpProviderClient{
		client: client,
	}, err
}

func (s *IpProviderClient) GetIP() (string, error) {
	return s.client.GetIP()
}

func (s *IpProviderClient) ShowIpProvider() string {
	return s.client.ShowIpProvider()
}
