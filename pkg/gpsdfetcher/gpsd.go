package gpsdfetcher

import (
	"context"
	"fmt"
	"time"

	"github.com/stratoberry/go-gpsd"
)

type GPSDFetcher struct {
	session *gpsd.Session
	Timeout time.Duration
}

// New creates and returns a new GPSFetcher instance
func NewGPSDFetcher(timeout time.Duration) *GPSDFetcher {
	return &GPSDFetcher{
		Timeout: timeout,
	}
}

// New creates and returns a new GPSFetcher instance
func NewWireGPSDFetcher() *GPSDFetcher {
	return &GPSDFetcher{}
}

// Connect to the gpsd service
func (g *GPSDFetcher) Connect() error {
	session, err := gpsd.Dial(gpsd.DefaultAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to GPSD: %w", err)
	}
	g.session = session
	return nil
}

// Fetch waits for a TPVReport with non-zero altitude and returns it
func (g *GPSDFetcher) Fetch() (*gpsd.TPVReport, error) {
	reportChan := make(chan *gpsd.TPVReport)
	ctx, cancel := context.WithCancel(context.Background()) // Create a context with cancellation

	defer cancel() // Ensure cancellation is called to clean up resources

	tpvFilter := func(r interface{}) {
		if tpv, ok := r.(*gpsd.TPVReport); ok && tpv.Alt != 0 {
			select {
			case reportChan <- tpv: // Attempt to send tpv
			case <-ctx.Done(): // If the context is cancelled, exit the filter function
				return
			}
		}
	}

	g.session.AddFilter("TPV", tpvFilter)
	g.session.Watch()

	select {
	case tpvReport := <-reportChan:
		return tpvReport, nil
	case <-time.After(g.Timeout):
		return nil, fmt.Errorf("timeout waiting for TPV report with non-zero altitude")
	}
	// Note: We're no longer closing the done channel here,
}

// Close cleanly closes the gpsd session
func (g *GPSDFetcher) Close() error {
	if g.session != nil {
		if err := g.session.Close(); err != nil {
			return fmt.Errorf("failed to close GPSD session: %w", err)
		}
	}
	return nil
}
