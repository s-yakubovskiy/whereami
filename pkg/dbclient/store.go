package dbclient

import (
	"fmt"
	"time"

	"github.com/s-yakubovskiy/whereami/pkg/contracts"
)

func (s *LocationKeeper) StoreLocation(location *contracts.Location) error {
	// Check if a record with the same query and city already exists
	exists, err := s.recordExists(location.IP, location.City)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("The database is already contains this record.")
	}

	// Get the current date in the desired format
	location.Date = time.Now().Format("2006-01-02 15:04:05")

	// Insert the new location
	stmt, err := s.db.Prepare(`
        INSERT INTO locations 
        (status, country, countryCode, region, regionName, city, zip, lat, lon, timezone, isp, org, asField, ip, date, vpn) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(location.Status, location.Country, location.CountryCode, location.Region,
		location.RegionName, location.City, location.Zip, location.Lat, location.Lon,
		location.Timezone, location.Isp, location.Org, location.As, location.IP, location.Date, location.Vpn)
	if err != nil {
		return err
	}

	return nil
}
