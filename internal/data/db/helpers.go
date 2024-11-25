package db

import (
	"database/sql"
	"log"
)

func (s *LocationKeeper) recordExists(query, city string) (bool, error) {
	var exists bool
	stmt, err := s.db.Prepare("SELECT EXISTS(SELECT 1 FROM locations WHERE ip = ? AND city = ? LIMIT 1)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(query, city).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *LocationKeeper) columnExists(columnName string) bool {
	query := `PRAGMA table_info(locations);`
	rows, err := s.db.Query(query)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	var cid int
	var name, colType string
	var notNull, pk int
	var dfltValue sql.NullString // Use sql.NullString for nullable columns

	for rows.Next() {
		if err := rows.Scan(&cid, &name, &colType, &notNull, &dfltValue, &pk); err != nil {
			log.Fatalf("Failed to read rows: %v", err)
		}
		if name == columnName {
			return true
		}
	}

	return false
}
