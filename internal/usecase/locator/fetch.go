package locator

// import (
// 	"context"
// 	"fmt"

// 	"github.com/s-yakubovskiy/whereami/internal/contracts"
// )

// func setupFetchRoutines(ctx context.Context, ip string, locationChan chan *contracts.Location, qualityChan chan *contracts.LocationScores, gpsReportChan chan *contracts.GPSReport, errorChan chan error, l *UseCase) {
// 	// Fetching data from IP API concurrently
// 	go func() {
// 		location, err := l.client.GetLocation(ip)
// 		if err != nil {
// 			errorChan <- fmt.Errorf("location error: %w", err)
// 			return
// 		}
// 		select {
// 		case locationChan <- location:
// 		case <-ctx.Done():
// 			errorChan <- fmt.Errorf("location fetch canceled")
// 		}
// 	}()

// 	// Fetch IP Quality scores concurrently
// 	// lCfg := whereami.NewConfig(cfg.ProviderConfigs.IpQualityScore.Enabled, ipLookup, cfg.GPSConfig.Enabled)

// 	if l.cfg.ProviderConfigs.IpQualityScore.Enabled {
// 		go func() {
// 			quality, err := l.client.AddIPQuality(ip)
// 			if err != nil {
// 				errorChan <- fmt.Errorf("quality error: %w", err)
// 				return
// 			}
// 			select {
// 			case qualityChan <- quality:
// 			case <-ctx.Done():
// 				errorChan <- fmt.Errorf("quality fetch canceled")
// 			}
// 		}()
// 	} else {
// 		close(qualityChan) // Close the channel if not used
// 	}

// 	// Fetch GPS report concurrently if enabled
// 	if l.cfg.GPSConfig.Enabled {
// 		go func() {
// 			report, err := l.FetchGPSReport()
// 			if err != nil || report == nil {
// 				errorChan <- fmt.Errorf("GPS error: %w", err)
// 				close(gpsReportChan) // Ensure to close the channel on failure
// 				return
// 			}
// 			select {
// 			case gpsReportChan <- contracts.GPSReportDTO(report):
// 			case <-ctx.Done():
// 				errorChan <- fmt.Errorf("GPS fetch canceled")
// 				close(gpsReportChan) // Ensure to close the channel if context is done
// 			}
// 		}()
// 	} else {
// 		close(gpsReportChan) // Close the channel if not used
// 	}
// }
