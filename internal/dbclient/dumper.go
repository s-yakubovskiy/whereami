package dbclient

import (
	"log"

	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

func (s *LocationKeeper) GetAllLocations() ([]contracts.Location, error) {
	rows, err := s.db.Query(`SELECT status, country, countryCode, region, regionName, city, zip, lat, lon, timezone, isp, org, asField, ip, date FROM locations`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var locations []contracts.Location
	for rows.Next() {
		var loc contracts.Location
		err = rows.Scan(&loc.Status, &loc.Country, &loc.CountryCode, &loc.Region, &loc.RegionName, &loc.City, &loc.Zip, &loc.Lat, &loc.Lon, &loc.Timezone, &loc.Isp, &loc.Org, &loc.As, &loc.IP, &loc.Date)
		if err != nil {
			log.Fatal(err)
		}
		locations = append(locations, loc)
	}
	return locations, nil
}

func (s *LocationKeeper) ImportLocations(locations []contracts.Location) error {
	for _, location := range locations {
		_, err := s.db.Exec(`INSERT INTO locations (status, country, countryCode, region, regionName, city, zip, lat, lon, timezone, isp, org, asField, ip, date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			location.Status, location.Country, location.CountryCode, location.Region, location.RegionName, location.City, location.Zip, location.Lat, location.Lon, location.Timezone, location.Isp, location.Org, location.As, location.IP, location.Date)
		if err != nil {
			log.Println("Failed to insert:", err)
		}
	}
	return nil
}
