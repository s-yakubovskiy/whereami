package dbclient

import (
	"log"

	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

// ShowLocations fetches a specified number of locations from the database.
func (s *LocationKeeper) ShowLocations(num int) ([]*contracts.Location, error) {
	var locations []*contracts.Location

	// Prepare the SQL query to fetch a limited number of locations.
	query := `SELECT ip, country, countryCode, region, regionCode, city, timezone, zip, flag, isp, asn, latitude, longitude, date, vpn FROM locations ORDER BY id DESC LIMIT ?`
	rows, err := s.db.Query(query, num)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over each row and scan the data into a Location struct.
	for rows.Next() {
		var loc contracts.Location
		err := rows.Scan(&loc.IP, &loc.Country, &loc.CountryCode,
			&loc.Region, &loc.RegionCode, &loc.City,
			&loc.Timezone, &loc.Zip, &loc.Flag,
			&loc.Isp, &loc.Asn, &loc.Latitude,
			&loc.Longitude, &loc.Date, &loc.Vpn)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		locations = append(locations, &loc)
	}

	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return locations, nil
}
