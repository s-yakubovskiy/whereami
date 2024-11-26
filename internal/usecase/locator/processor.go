package locator

import (
	"context"
	"fmt"

	"github.com/s-yakubovskiy/whereami/internal/entity"
)

func (uc *UseCase) setupFetchRoutines(ctx context.Context, ip string, locationChan chan *entity.Location, qualityChan chan *entity.LocationScores, errorChan chan error) {
	// Fetching data from IP API concurrently
	go func() {
		location, err := uc.ipInfoRepo.LookupIpInfo(ip)
		if err != nil {
			errorChan <- uc.log.ErrorfNew("location error: %w", err)
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

	if uc.cfg.ProviderConfigs.IpQualityScore.Enabled {
		go func() {
			quality, err := uc.ipQualityScoreRepo.LookupIpQualityScore(ip)
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
}

func (uc *UseCase) collectResults(ctx context.Context, locationChan chan *entity.Location, qualityChan chan *entity.LocationScores, errorChan chan error) (*entity.Location, *entity.LocationScores) {
	var location *entity.Location
	var quality *entity.LocationScores

	completed := 0
	total := 2

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
		case err := <-errorChan:
			uc.log.Errorf("Operation error: " + err.Error())
		case <-ctx.Done():
			uc.log.Warnf("Timeout or context cancellation / " + ctx.Err().Error())
			return location, quality
		}
	}
	return location, quality
}
