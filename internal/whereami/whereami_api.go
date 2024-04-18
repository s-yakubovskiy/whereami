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
	if l.cfg.IpQuality {
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
	if l.cfg.GPS {
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

// func (l *Locator) ShowFull() {
// 	var ip string
// 	var err error

// 	if l.cfg.IP == "" {
// 		ip, err = l.client.GetIP()
// 		if err != nil {
// 			common.Errorln(err.Error())
// 			return
// 		}
// 	} else {
// 		ip = l.cfg.IP
// 	}

// 	// Create channels for concurrent fetching
// 	locationChan := make(chan *contracts.Location, 1)
// 	qualityChan := make(chan *contracts.LocationScores, 1)
// 	gpsReportChan := make(chan *contracts.GPSReport, 1)
// 	errorChan := make(chan error, 3) // to handle errors from goroutines

// 	// Fetching data from IP API concurrently
// 	go func() {
// 		location, err := l.client.GetLocation(ip)
// 		if err != nil {

// 			errorChan <- err
// 			close(locationChan)
// 			return
// 		}
// 		locationChan <- location
// 	}()

// 	// Fetch IP Quality scores concurrently
// 	if l.cfg.IpQuality {
// 		go func() {
// 			quality, err := l.client.AddIPQuality(ip)
// 			if err != nil {
// 				errorChan <- err
// 				close(qualityChan)
// 				return
// 			}
// 			qualityChan <- quality
// 		}()
// 	} else {
// 		close(qualityChan) // Close the channel if not used
// 	}

// 	// Fetch GPS report concurrently if enabled
// 	if l.cfg.GPS {
// 		go func() {
// 			report, err := l.FetchGPSReport()
// 			if err != nil || report == nil {
// 				// log.Fatal("err gps", err, "report", report)
// 				errorChan <- err
// 				close(gpsReportChan)
// 				return
// 			}
// 			gpsReportChan <- contracts.GPSReportDTO(report)
// 		}()
// 	} else {
// 		close(gpsReportChan) // Close the channel if not used
// 	}

// 	// Wait for all results
// 	location := <-locationChan
// 	quality := <-qualityChan
// 	gpsReport := <-gpsReportChan

// 	// Check for any errors from goroutines
// 	close(errorChan)
// 	for err := range errorChan {
// 		common.Errorln(err.Error())
// 		// just inform about errors
// 		// return
// 	}

// 	// VPN checking is synchronous
// 	vpninterfaces, err := l.dbclient.GetVPNInterfaces()
// 	if err != nil {
// 		common.Warnln(err.Error())
// 	}

// 	vpn, err := l.client.GetVPN(vpninterfaces)
// 	if err != nil {
// 		common.Warnln(err.Error())
// 	}

// 	if vpn {
// 		location.Vpn = true
// 	}

// 	// Combine all data into the final Location struct
// 	if quality != nil {
// 		location.Scores = *quality
// 	}
// 	if gpsReport != nil {
// 		location.Gps = *gpsReport
// 	}
// 	location.Comment += ". Using public ip provider: " + l.client.ShowIpProvider()

// 	location.Output(categories, orderedCategories)
// }
