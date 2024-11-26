package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// LocationKeeper implements the DBInterface for SQLite
type LocationKeeper struct {
	db *sql.DB
	// log
}

func NewLocationKeeper(databasePath string) (*LocationKeeper, error) {
	if databasePath[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		databasePath = filepath.Join(homeDir, databasePath[1:])
	}

	// Create the directory if it doesn't exist
	dir := filepath.Dir(databasePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	// Open the SQLite database
	db, err := sql.Open("sqlite", databasePath)
	if err != nil {
		return nil, err
	}

	return &LocationKeeper{db: db}, nil
}
