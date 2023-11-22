package dbclient

import (
	"log"
)

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

	// add vpn_interfaces table
	if err := s.createVPNTable(); err != nil {
		return err
	}

	// alter table with vpn column
	if err := s.addVPNColumToTable(); err != nil {
		return err
	}

	return nil
}

func (s *LocationKeeper) createVPNTable() error {
	// vpn_interfaces table
	createVPNInterfacesTableSQL := `
    CREATE TABLE IF NOT EXISTS vpn_interfaces (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        interface_name TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	stmt, err := s.db.Prepare(createVPNInterfacesTableSQL)
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

// migration with add vpn column
func (s *LocationKeeper) addVPNColumToTable() error {
	if !s.columnExists("vpn") {
		statement := `ALTER TABLE locations ADD COLUMN vpn INTEGER;`
		_, err := s.db.Exec(statement)
		if err != nil {
			log.Printf("Error adding column 'vpn': %v", err)
			return err
		}
		log.Println("Column 'vpn' added successfully.")
	} else {
		log.Println("Column 'vpn' already exists.")
	}
	return nil
}
