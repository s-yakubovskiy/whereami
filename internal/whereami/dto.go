package whereami

import (
	"fmt"

	"github.com/stratoberry/go-gpsd"
)

func GPSReportDTO(report *gpsd.TPVReport) *GPSReport {
	return &GPSReport{
		Latitude:  report.Lat,
		Longitude: report.Lon,
		Altitude:  report.Alt,
		Url:       fmt.Sprintf("https://www.google.com/maps?q=%f,%f", report.Lat, report.Lon),
		// url: contracts.CreateMapLocation(report.Lat, report.Lon)
	}
}
