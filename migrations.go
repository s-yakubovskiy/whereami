package migrations

import (
	"database/sql"
	"embed"
	"log"
	"os"
	"path/filepath"

	"github.com/pressly/goose"
)

//go:embed db/migrations/*.sql

var embedMigrations embed.FS //

// GooseWorker implements the DBInterface for SQLite
type GooseWorker struct {
	db *sql.DB
}

func NewGooseWorker(dataSourceName string) (*GooseWorker, error) {
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

	return &GooseWorker{db: db}, nil
}

func (s *GooseWorker) InitDB() error {
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
