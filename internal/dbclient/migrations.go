package dbclient

import (
	"log"

	"github.com/pressly/goose"
)

func (s *LocationKeeper) InitDB() error {
	// Run migrations using Goose
	err := goose.SetDialect("sqlite3")
	if err != nil {
		return err
	}

	err = goose.Up(s.db, "db/migrations")
	if err != nil {
		log.Printf("Goose migration failed: %v", err)
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
