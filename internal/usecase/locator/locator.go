package locator

import (
	"context"
	"time"

	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/pkg/shudralogs"

	"github.com/s-yakubovskiy/whereami/internal/entity"
	"github.com/s-yakubovskiy/whereami/internal/metrics"
)

var ASYNC_TIMEOUT = 35 * time.Second

var (
	categories = map[string][]string{
		"Geographical Information": {
			"country", "countryCode", "region", "regionCode",
			"city", "timezone", "zip", "latitude", "longitude",
		},
		"Network Information": {
			"ip", "isp", "asn", "flag",
		},
		"Security Assessments": {
			"vpnInterface", "scores",
		},
		// "GPS": {
		// 	"gps",
		// },
		"Miscellaneous": {
			"map", "date", "comment",
		},
	}

	orderedCategories = []string{
		"Network Information",
		"Geographical Information",
		"Security Assessments",
		"GPS",
		"Miscellaneous",
	}
)

type UseCase struct {
	cfg                *config.AppConfig
	publicIpRepo       PublicIpRepo
	ipInfoRepo         IpInfoRepo
	ipQualityScoreRepo IpQualityScoreRepo
	log                shudralogs.Logger
	m                  metrics.Metrics
}

type IpInfoRepo interface {
	LookupIpInfo(string) (*entity.Location, error)
}

type IpQualityScoreRepo interface {
	LookupIpQualityScore(string) (*entity.LocationScores, error)
}

type PublicIpRepo interface {
	ShowIpProvider() string
	GetIP() (string, error)
}

func NewLocatorUserCase(log shudralogs.Logger, cfg *config.AppConfig, locationRepo PublicIpRepo, ipInfoRepo IpInfoRepo, ipQualityRepo IpQualityScoreRepo, m metrics.Metrics) *UseCase {
	// Register metrics specific to this use case
	m.RegisterCounter("show_location_count", "Counts custom ShowLocation executions", []string{"status"})
	m.RegisterHistogram("show_location_latency", "Tracks custom ShowLocation latencies", []string{"task_type"})
	return &UseCase{
		cfg:                cfg,
		log:                log,
		publicIpRepo:       locationRepo,
		ipInfoRepo:         ipInfoRepo,
		ipQualityScoreRepo: ipQualityRepo,
		m:                  m,
	}
}

func (uc *UseCase) ShowLocation(ctx context.Context) (*entity.Location, error) {
	start := time.Now()
	location, err := uc.getLocation(ctx)
	if err != nil {
		return nil, err
	}
	uc.Output(location, categories, orderedCategories)
	uc.m.IncrementCounter("show_location_count", map[string]string{"status": "success"})
	uc.m.ObserveLatency("show_location_latency", time.Since(start), map[string]string{"task_type": "ShowLocation"})
	return location, nil
}

func (uc *UseCase) GetLocation(ctx context.Context) (*entity.Location, error) {
	return uc.getLocation(ctx)
}

func (uc *UseCase) getLocation(ctx context.Context) (*entity.Location, error) {
	var ip string
	var err error

	uc.log.Debugf("UseCase ShowLocation, public ip %s", uc.cfg.PublicIP)

	if uc.cfg.PublicIP == "" {
		ip, err = uc.publicIpRepo.GetIP()
		if err != nil {
			uc.log.Errorf("Error fetching IP: %v", err.Error())
		}
	} else {
		ip = uc.cfg.PublicIP
	}

	// Create channels for concurrent fetching
	locationChan := make(chan *entity.Location, 1)
	qualityChan := make(chan *entity.LocationScores, 1)
	errorChan := make(chan error, 3) // to handle errors from goroutines
	uc.setupFetchRoutines(
		ctx, ip, locationChan, qualityChan, errorChan,
	)

	location, quality := uc.collectResults(ctx, locationChan, qualityChan, errorChan)

	// Combine all data into the final Location struct
	if quality != nil {
		location.Scores = *quality
	}
	location.Comment += ". Using public ip provider: <tbd>"
	return location, nil
}
