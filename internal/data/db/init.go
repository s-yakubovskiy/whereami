package db

import (
	"context"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed db/migrations/*.sql

var embedMigrations embed.FS //

func (s *LocationKeeper) InitDb(ctx context.Context) error {
	// Run migrations using Goose
	err := goose.SetDialect("sqlite3")
	goose.SetBaseFS(embedMigrations)
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
