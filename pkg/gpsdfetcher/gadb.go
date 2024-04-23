package gpsdfetcher

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/electricbubble/gadb"
	"github.com/stratoberry/go-gpsd"
)

// check compat
var _ GPSInterface = &GPSAdbFetcher{}

type GPSAdbFetcher struct {
	Timeout time.Duration
	dev     *gadb.Device
}

func NewGPSAdbFetcher() *GPSAdbFetcher {
	return &GPSAdbFetcher{}
}

func (s *GPSAdbFetcher) Close() error {
	return nil
}

func (s *GPSAdbFetcher) Connect() error {
	adbClient, err := gadb.NewClient()
	checkErr(err, "fail to connect adb server")

	devices, err := adbClient.DeviceList()
	checkErr(err)

	if len(devices) == 0 {
		log.Fatalln("list of devices is empty")
	}

	s.dev = &devices[0]

	// userHomeDir, _ := os.UserHomeDir()

	// fmt.Printf("%+v\n", s.dev)
	// state, err := s.dev.State()
	// fmt.Printf("\n%+v\nerr: %+v\n", state, err)

	return nil
}

func (s *GPSAdbFetcher) Fetch() (*gpsd.TPVReport, error) {
	shellOutput, err := s.dev.RunShellCommand("dumpsys location")
	if err != nil {
		log.Fatalf("error running shell command: %v", err)
	}

	return parseLocation(shellOutput)
}

func checkErr(err error, msg ...string) {
	if err == nil {
		return
	}

	var output string
	if len(msg) != 0 {
		output = msg[0] + " "
	}
	output += err.Error()
	log.Fatalln(output)
}

func parseLocation(output string) (*gpsd.TPVReport, error) {
	// Regular expression to match the location data
	re := regexp.MustCompile(`last location=Location\[(?P<Type>\w+) (?P<Latitude>\d+\.\d+),(?P<Longitude>\d+\.\d+) hAcc=(?P<HorizontalAcc>\d+\.\d+) et=(?P<ElapsedTime>.+?) alt=(?P<Altitude>\d+\.\d+) vAcc=(?P<VerticalAcc>\d+\.\d+) mslAlt=(?P<MslAltitude>\d+\.\d+) mslAltAcc=(?P<MslAltAcc>\d+\.\d+) {Bundle\[(?P<Extras>.+?)\]}\]`)

	// Find match
	match := re.FindStringSubmatch(output)
	if match == nil {
		return nil, fmt.Errorf("no location data found")
	}

	// Map the names to captured groups
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	// Convert captured strings to appropriate types
	latitude, _ := strconv.ParseFloat(result["Latitude"], 64)
	longitude, _ := strconv.ParseFloat(result["Longitude"], 64)
	// horizontalAcc, _ := strconv.ParseFloat(result["HorizontalAcc"], 64)
	altitude, _ := strconv.ParseFloat(result["Altitude"], 64)
	// verticalAcc, _ := strconv.ParseFloat(result["VerticalAcc"], 64)
	// mslAltitude, _ := strconv.ParseFloat(result["MslAltitude"], 64)
	// mslAltAcc, _ := strconv.ParseFloat(result["MslAltAcc"], 64)

	extras := make(map[string]float64)
	extrasStr := strings.Trim(result["Extras"], "{}")
	for _, pair := range strings.Split(extrasStr, ",") {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			val, _ := strconv.ParseFloat(kv[1], 64)
			extras[kv[0]] = val
		}
	}

	// Create the Location struct
	return &gpsd.TPVReport{
		Class: result["Type"],
		Lat:   latitude,
		Lon:   longitude,
		Alt:   altitude,
		// Type:          result["Type"],
		// Latitude:      latitude,
		// Longitude:     longitude,
		// HorizontalAcc: horizontalAcc,
		// ElapsedTime:   result["ElapsedTime"],
		// Altitude:      altitude,
		// VerticalAcc:   verticalAcc,
		// MslAltitude:   mslAltitude,
		// MslAltAcc:     mslAltAcc,
		// Extra:         extras,
	}, nil
}
