package whereami

import (
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
			"vpn", "scores",
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

	// Create channels for concurrent fetching
	locationChan := make(chan *contracts.Location, 1)
	qualityChan := make(chan *contracts.LocationScores, 1)
	gpsReportChan := make(chan *contracts.GPSReport, 1)
	errorChan := make(chan error, 3) // to handle errors from goroutines

	// Fetching data from IP API concurrently
	go func() {
		location, err := l.client.GetLocation(ip)
		if err != nil {
			errorChan <- err
			return
		}
		locationChan <- location
	}()

	// Fetch IP Quality scores concurrently
	if l.cfg.IpQuality {
		go func() {
			quality, err := l.client.AddIPQuality(ip)
			if err != nil {
				errorChan <- err
				return
			}
			qualityChan <- quality
		}()
	} else {
		close(qualityChan) // Close the channel if not used
	}

	// Fetch GPS report concurrently if enabled
	if l.cfg.GPS {
		go func() {
			report, err := l.FetchGPSReport()
			if err != nil {
				errorChan <- err
				return
			}
			gpsReportChan <- contracts.GPSReportDTO(report)
		}()
	} else {
		close(gpsReportChan) // Close the channel if not used
	}

	// Wait for all results
	location := <-locationChan
	quality := <-qualityChan
	gpsReport := <-gpsReportChan

	// Check for any errors from goroutines
	close(errorChan)
	for err := range errorChan {
		common.Errorln(err.Error())
		return
	}

	// VPN checking is synchronous
	vpninterfaces, err := l.dbclient.GetVPNInterfaces()
	if err != nil {
		common.Warnln(err.Error())
	}

	vpn, err := l.client.GetVPN(vpninterfaces)
	if err != nil {
		common.Warnln(err.Error())
	}

	if vpn {
		location.Vpn = true
	}

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
