package migration

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"path/filepath"
	"runtime"
)

func NewPgMigrator(db *sql.DB) (*migrate.Migrate, error) {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("failed to get path")
	}

	sourceUrl := "file://" + filepath.Dir(path) + "/pg"

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatalf("failed to create migrator driver: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(sourceUrl, "postgres", driver)

	return m, err
}
