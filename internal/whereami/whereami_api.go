package whereami

import (
	"context"
	"fmt"

	"github.com/s-yakubovskiy/whereami/internal/common"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
	"github.com/stratoberry/go-gpsd"
)

func (l *Locator) FetchGPSReport() (*gpsd.TPVReport, error) {
	if err := l.gps.Connect(); err != nil {
		return nil, fmt.Errorf("Error connecting to GPSD: %v", err)
	}
	defer l.gps.Close()

	tpvReport, err := l.gps.Fetch()
	if err != nil {
		return nil, fmt.Errorf("Error fetching TPV report: %v", err)
	}

	return tpvReport, nil
}

func setupFetchRoutines(ctx context.Context, ip string, locationChan chan *contracts.Location, qualityChan chan *contracts.LocationScores, gpsReportChan chan *contracts.GPSReport, errorChan chan error, l *Locator) {
	// Fetching data from IP API concurrently
	go func() {
		location, err := l.client.GetLocation(ip)
		if err != nil {
			errorChan <- fmt.Errorf("location error: %w", err)
			return
		}
		select {
		case locationChan <- location:
		case <-ctx.Done():
			errorChan <- fmt.Errorf("location fetch canceled")
		}
	}()

	// Fetch IP Quality scores concurrently
	// lCfg := whereami.NewConfig(cfg.ProviderConfigs.IpQualityScore.Enabled, ipLookup, cfg.GPSConfig.Enabled)

	if l.cfg.ProviderConfigs.IpQualityScore.Enabled {
		go func() {
			quality, err := l.client.AddIPQuality(ip)
			if err != nil {
				errorChan <- fmt.Errorf("quality error: %w", err)
				return
			}
			select {
			case qualityChan <- quality:
			case <-ctx.Done():
				errorChan <- fmt.Errorf("quality fetch canceled")
			}
		}()
	} else {
		close(qualityChan) // Close the channel if not used
	}

	// Fetch GPS report concurrently if enabled
	if l.cfg.GPSConfig.Enabled {
		go func() {
			report, err := l.FetchGPSReport()
			if err != nil || report == nil {
				errorChan <- fmt.Errorf("GPS error: %w", err)
				close(gpsReportChan) // Ensure to close the channel on failure
				return
			}
			select {
			case gpsReportChan <- contracts.GPSReportDTO(report):
			case <-ctx.Done():
				errorChan <- fmt.Errorf("GPS fetch canceled")
				close(gpsReportChan) // Ensure to close the channel if context is done
			}
		}()
	} else {
		close(gpsReportChan) // Close the channel if not used
	}
}

func collectResults(ctx context.Context, locationChan chan *contracts.Location, qualityChan chan *contracts.LocationScores, gpsReportChan chan *contracts.GPSReport, errorChan chan error, gpsEnabled bool) (*contracts.Location, *contracts.LocationScores, *contracts.GPSReport) {
	var location *contracts.Location
	var quality *contracts.LocationScores
	var gpsReport *contracts.GPSReport

	completed := 0
	total := 2 // default for when GPS is not enabled
	if gpsEnabled {
		total = 3 // adjust total if GPS is enabled
	}

	for completed < total {
		select {
		case loc, ok := <-locationChan:
			if ok {
				location = loc
				completed++
			}
		case qual, ok := <-qualityChan:
			if ok {
				quality = qual
				completed++
			}
		case gps, ok := <-gpsReportChan:
			if ok {
				gpsReport = gps
				completed++
			} else if gpsEnabled {
				// GPS channel closed due to error, proceed without GPS data
				completed++
			}
		case err := <-errorChan:
			common.Errorln("Operation error: " + err.Error())
		case <-ctx.Done():
			common.Warnln("Timeout or context cancellation / " + ctx.Err().Error())
			return location, quality, gpsReport // Return whatever was fetched
		}
	}
	return location, quality, gpsReport
}
