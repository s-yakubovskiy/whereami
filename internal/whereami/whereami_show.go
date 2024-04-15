package whereami

import (
	"context"
	"time"

	"github.com/s-yakubovskiy/whereami/internal/common"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

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
		"GPS": {
			"gps",
		},
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

func (l *Locator) Show() {
	// NOTE: right now no difference between full and short `show` output
	l.ShowFull()
}

func (l *Locator) ShowFull() {
	var ip string
	var err error

	if l.cfg.IP == "" {
		ip, err = l.client.GetIP()
		if err != nil {
			common.Errorln(err.Error())
			return
		}
	} else {
		ip = l.cfg.IP
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel() // Ensures all paths cancel the context to prevent leaks

	// Create channels for concurrent fetching
	locationChan := make(chan *contracts.Location, 1)
	qualityChan := make(chan *contracts.LocationScores, 1)
	gpsReportChan := make(chan *contracts.GPSReport, 1)
	errorChan := make(chan error, 3) // to handle errors from goroutines

	// Setting up fetch routines
	setupFetchRoutines(ctx, ip, locationChan, qualityChan, gpsReportChan, errorChan, l)

	// Collect results and handle possible timeouts
	location, quality, gpsReport := collectResults(ctx, locationChan, qualityChan, gpsReportChan, errorChan)

	// Combine all data into the final Location struct
	if quality != nil {
		location.Scores = *quality
	}
	if gpsReport != nil {
		location.Gps = *gpsReport
	}
	location.Comment += ". Using public ip provider: " + l.client.ShowIpProvider()
	location.Output(categories, orderedCategories)
}
