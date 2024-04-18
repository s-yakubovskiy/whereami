package dbclient

import (
	"fmt"
	"time"

	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

func (s *LocationKeeper) StoreLocation(location *contracts.Location) error {
	// Check if a record with the same IP and city already exists
	exists, err := s.recordExists(location.IP, location.City)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("The database already contains this record.")
	}

	// Get the current date in the desired format
	location.Date = time.Now().Format("2006-01-02 15:04:05")

	// Begin a transaction
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Insert the new location
	stmt, err := tx.Prepare(`
        INSERT INTO locations 
        (ip, country, countryCode, region, regionCode, city, timezone, zip, flag, isp, asn, latitude, longitude, date, vpn, gps_latitude, gps_longitude, gps_altitude, gps_url)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		location.IP, location.Country, location.CountryCode, location.Region, location.RegionCode,
		location.City, location.Timezone, location.Zip, location.Flag, location.Isp, location.Asn,
		location.Latitude, location.Longitude, location.Date, location.Vpn,
		location.Gps.Latitude, location.Gps.Longitude, location.Gps.Altitude, location.Gps.Url)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Get the last inserted ID
	lastID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert scores data into location_scores
	stmt, err = tx.Prepare(`
        INSERT INTO location_scores
        (location_id, fraud_score, host, proxy, vpn, tor, is_crawler, recent_abuse, bot_status)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(lastID, location.Scores.FraudScore, location.Scores.Host, location.Scores.Proxy, location.Scores.VPN, location.Scores.Tor, location.Scores.IsCrawler, location.Scores.RecentAbuse, location.Scores.BotStatus)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
