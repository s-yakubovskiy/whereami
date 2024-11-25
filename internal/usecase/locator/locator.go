package locator

import (
	"context"
	"time"

	"github.com/s-yakubovskiy/whereami/internal/config"
	// "github.com/s-yakubovskiy/whereami/internal/contracts"

	"github.com/s-yakubovskiy/whereami/internal/entity"
	"github.com/s-yakubovskiy/whereami/internal/logging"
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
	log                logging.Logger
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

func NewLocatorUserCase(log logging.Logger, cfg *config.AppConfig, locationRepo PublicIpRepo, ipInfoRepo IpInfoRepo, ipQualityRepo IpQualityScoreRepo) *UseCase {
	return &UseCase{
		cfg:                cfg,
		log:                log,
		publicIpRepo:       locationRepo,
		ipInfoRepo:         ipInfoRepo,
		ipQualityScoreRepo: ipQualityRepo,
	}
}

func (uc *UseCase) ShowLocation(ctx context.Context) (*entity.Location, error) {
	location, err := uc.getLocation(ctx)
	if err != nil {
		return nil, err
	}
	uc.Output(location, categories, orderedCategories)
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
