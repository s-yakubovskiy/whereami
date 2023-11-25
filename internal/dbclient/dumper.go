package dbclient

import (
	"log"

	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

func (s *LocationKeeper) GetAllLocations() ([]contracts.Location, error) {
	rows, err := s.db.Query(`SELECT ip, country, countryCode, region, regionCode, city, timezone, zip, flag, emojiFlag, isp, org, asn, latitude, longitude, date, vpn FROM locations`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var locations []contracts.Location
	for rows.Next() {
		var loc contracts.Location
		err = rows.Scan(&loc.IP, &loc.Country, &loc.CountryCode, &loc.Region, &loc.RegionCode, &loc.City, &loc.Timezone, &loc.Zip, &loc.Flag, &loc.EmojiFlag, &loc.Isp, &loc.Org, &loc.Asn, &loc.Latitude, &loc.Longitude, &loc.Date, &loc.Vpn)
		if err != nil {
			log.Fatal(err)
		}
		locations = append(locations, loc)
	}
	return locations, nil
}

func (s *LocationKeeper) ImportLocations(locations []contracts.Location) error {
	for _, location := range locations {
		log.Fatalf("%+v\n", location)
		_, err := s.db.Exec(`INSERT INTO locations (ip, country, countryCode, region, regionCode, city, timezone, zip, flag, emojiFlag, isp, org, asn, latitude, longitude, date, vpn) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			location.IP, location.Country, location.CountryCode, location.Region, location.RegionCode, location.City, location.Timezone, location.Zip, location.Flag, location.EmojiFlag, location.Isp, location.Org, location.Asn, location.Latitude, location.Longitude, location.Date, location.Vpn)
		if err != nil {
			log.Println("Failed to insert:", err)
		}
	}
	return nil
}

// ip, country, countryCode, region, regionCode, city, timezone, zip, postal, flag, emojiFlag, isp, org, asn, latitude, longitude, date, vpn
