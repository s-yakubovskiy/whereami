package apimanager

import (
	ipdata "github.com/ipdata/go"
	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
	"github.com/s-yakubovskiy/whereami/pkg/ipapi"
	"github.com/s-yakubovskiy/whereami/pkg/ipqualityscore"
)

var _ IPQualityInterface = &IpQualityScoreClient{}

type IPQualityInterface interface {
	AddIPQuality(*contracts.Location, string) (*contracts.Location, error)
}

// var _ v1.CloudSQLServiceServer = &CloudSQLService{}

// IpApi client initialization
type IpApiClient struct {
	client *ipapi.IpApi
}

func NewIpApiClient(cfg config.ProviderConfig) (*IpApiClient, error) {
	client, err := ipapi.NewIpApi(cfg.URL, cfg.APIKey)
	return &IpApiClient{
		client: client,
	}, err
}

func (l *IpApiClient) GetLocation(ip string) (*contracts.Location, error) {
	data, err := l.client.Lookup(ip)
	if err != nil {
		return nil, err
	}
	return ConvertIpApiToLocation(data)
}

// IpData client initialization
type IpDataClient struct {
	client *ipdata.Client
}

func NewIpDataClient(cfg config.ProviderConfig) (*IpDataClient, error) {
	client, err := ipdata.NewClient(cfg.APIKey)
	return &IpDataClient{
		client: &client,
	}, err
}

func (l *IpDataClient) GetLocation(ip string) (*contracts.Location, error) {
	data, err := l.client.Lookup(ip)
	if err != nil {
		return nil, err
	}
	return ConvertIpDataToLocation(data)
}

// IpQualityScore client initialization
type IpQualityScoreClient struct {
	client *ipqualityscore.IpQualityScore
}

func NewIpQualityScoreClient(cfg config.ProviderConfig) (*IpQualityScoreClient, error) {
	client, err := ipqualityscore.NewIpQualityScore(cfg.URL, cfg.APIKey)
	return &IpQualityScoreClient{
		client: client,
	}, err
}

func (l *IpQualityScoreClient) AddIPQuality(location *contracts.Location, ip string) (*contracts.Location, error) {
	qualityScore, err := l.client.Lookup(ip)
	if err != nil {
		return location, err
	}

	return EnrichLocationWithQualityScore(location, qualityScore)
}
