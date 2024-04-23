package gpsdfetcher

import "github.com/stratoberry/go-gpsd"

type GPSInterface interface {
	Connect() error
	Close() error
	Fetch() (*gpsd.TPVReport, error)
}
