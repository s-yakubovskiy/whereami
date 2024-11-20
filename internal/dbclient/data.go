package dbclient

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/s-yakubovskiy/whereami/config"
)

// LocationKeeper implements the DBInterface for SQLite
type LocationKeeper struct {
	db *sql.DB
}

// NewSQLiteDB creates a new instance of SQLiteDB
func NewSQLiteDB(cfg config.Database) (*LocationKeeper, error) {
	// Expand the '~' to the user's home directory
	if cfg.Path[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		cfg.Path = filepath.Join(homeDir, cfg.Path[1:])
	}

	// Create the directory if it doesn't exist
	dir := filepath.Dir(cfg.Path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	// Open the SQLite database
	db, err := sql.Open("sqlite3", cfg.Path)
	if err != nil {
		return nil, err
	}

	return &LocationKeeper{db: db}, nil
}
