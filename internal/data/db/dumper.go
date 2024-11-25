package db

import (
	"log"

	"github.com/s-yakubovskiy/whereami/internal/entity"
)

func (s *LocationKeeper) GetAllLocations() ([]entity.Location, error) {
	query := `SELECT l.ip, l.country, l.countryCode, l.region, l.regionCode, l.city, l.timezone, l.zip, l.flag, l.isp, l.asn, l.latitude, l.longitude, l.date, l.vpn,
                     s.fraud_score, s.host, s.proxy, s.vpn, s.tor, s.is_crawler, s.recent_abuse, s.bot_status
              FROM locations l
              LEFT JOIN location_scores s ON l.id = s.location_id`

	rows, err := s.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var locations []entity.Location
	for rows.Next() {
		var loc entity.Location
		// Update the Scan method to include new fields
		err = rows.Scan(&loc.IP, &loc.Country, &loc.CountryCode, &loc.Region, &loc.RegionCode, &loc.City, &loc.Timezone, &loc.Zip, &loc.Flag, &loc.Isp, &loc.Asn, &loc.Latitude, &loc.Longitude, &loc.Date, &loc.Vpn,
			&loc.Scores.FraudScore, &loc.Scores.Host, &loc.Scores.Proxy, &loc.Scores.VPN, &loc.Scores.Tor, &loc.Scores.IsCrawler, &loc.Scores.RecentAbuse, &loc.Scores.BotStatus)
		if err != nil {
			log.Fatal(err)
		}
		locations = append(locations, loc)
	}
	return locations, nil
}

func (s *LocationKeeper) ImportLocations(locations []entity.Location) error {
	for _, location := range locations {
		tx, err := s.db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		// Insert into locations table
		res, err := tx.Exec(`INSERT INTO locations (ip, country, countryCode, region, regionCode, city, timezone, zip, flag, isp, asn, latitude, longitude, date, vpn) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			location.IP, location.Country, location.CountryCode, location.Region, location.RegionCode, location.City, location.Timezone, location.Zip, location.Flag, location.Isp, location.Asn, location.Latitude, location.Longitude, location.Date, location.Vpn)
		if err != nil {
			tx.Rollback()
			log.Println("Failed to insert into locations:", err)
			continue
		}

		// Get last insert ID
		lastID, err := res.LastInsertId()
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		// Insert into location_scores table
		_, err = tx.Exec(`INSERT INTO location_scores (location_id, fraud_score, host, proxy, vpn, tor, is_crawler, recent_abuse, bot_status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			lastID, location.Scores.FraudScore, location.Scores.Host, location.Scores.Proxy, location.Scores.VPN, location.Scores.Tor, location.Scores.IsCrawler, location.Scores.RecentAbuse, location.Scores.BotStatus)
		if err != nil {
			tx.Rollback()
			log.Println("Failed to insert into location_scores:", err)
			continue
		}

		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
