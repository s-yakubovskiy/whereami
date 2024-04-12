package whereami

import (
	"fmt"

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
