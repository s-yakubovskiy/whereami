package gpsdfetcher

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/stratoberry/go-gpsd"
)

var GPSD_DUMP_FILE_PATH = "/opt/gps/data/gpsdata.json"

type GPSDFileFetcher struct {
	Timeout time.Duration
}

// New creates and returns a new GPSFetcher instance
func NewGPSDFileFetcher(timeout time.Duration) *GPSDFileFetcher {
	return &GPSDFileFetcher{
		Timeout: timeout,
	}
}

// Connect to the gpsd dump file
func (g *GPSDFileFetcher) Connect() error {
	return nil
}

// Fetch data from dump json /opt/gps/data/gpsdata.json
func (g *GPSDFileFetcher) Fetch() (*gpsd.TPVReport, error) {
	data, err := os.ReadFile(GPSD_DUMP_FILE_PATH)
	if err != nil {
		return nil, fmt.Errorf("failed to read GPS data file: %w", err)
	}

	var tpvReport gpsd.TPVReport
	err = json.Unmarshal(data, &tpvReport)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal TPV report: %w", err)
	}

	// Optional: Validate the data if needed
	// if tpvReport.Mode == 0 { // Mode 0 might indicate no valid data
	// 	return nil, fmt.Errorf("invalid GPS data: no fix")
	// }

	return &tpvReport, nil
}

// Close cleanly closes the gpsd session
func (g *GPSDFileFetcher) Close() error {
	return nil
}
