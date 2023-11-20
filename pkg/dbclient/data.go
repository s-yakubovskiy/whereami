package dbclient

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// LocationKeeper implements the DBInterface for SQLite
type LocationKeeper struct {
	db *sql.DB
}

// NewSQLiteDB creates a new instance of SQLiteDB
func NewSQLiteDB(dataSourceName string) (*LocationKeeper, error) {
	// Expand the '~' to the user's home directory
	if dataSourceName[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		dataSourceName = filepath.Join(homeDir, dataSourceName[1:])
	}

	// Create the directory if it doesn't exist
	dir := filepath.Dir(dataSourceName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	// Open the SQLite database
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &LocationKeeper{db: db}, nil
}
