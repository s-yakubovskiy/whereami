package dbclient

func (s *LocationKeeper) InitDB() error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS locations (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        status TEXT,
        country TEXT,
        countryCode TEXT,
        region TEXT,
        regionName TEXT,
        city TEXT,
        zip TEXT,
        lat REAL,
        lon REAL,
        timezone TEXT,
        isp TEXT,
        org TEXT,
        asField TEXT,
        ip TEXT,
        date TEXT
    );`

	stmt, err := s.db.Prepare(createTableSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
