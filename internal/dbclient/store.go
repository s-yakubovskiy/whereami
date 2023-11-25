package dbclient

import (
	"fmt"
	"time"

	"github.com/s-yakubovskiy/whereami/internal/contracts"
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
				(ip, country, countryCode, region, regionCode, city, timezone, zip, flag, emojiFlag, isp, org, asn, latitude, longitude, date, vpn)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(location.IP, location.Country, location.CountryCode, location.Region, location.RegionCode, location.City, location.Timezone, location.Zip, location.Flag, location.EmojiFlag, location.Isp, location.Org, location.Asn, location.Latitude, location.Longitude, location.Date, location.Vpn)
	if err != nil {
		return err
	}

	return nil
}
