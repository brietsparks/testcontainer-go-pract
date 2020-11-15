package pg

import (
	"context"
	"database/sql"
	"log"
	"pract-testcontainers/db/pg"
	"testing"
)

var db *sql.DB

func TestMain(m *testing.M) {
	ctx := context.Background()

	container, database, err := pg.CreateTestContainer(ctx, "testdb")
	if err != nil {
		log.Fatal(err)
	}

	mig, err := pg.NewPgMigrator(database)
	if err != nil {
		log.Fatal(err)
	}

	err = mig.Up()
	if err != nil {
		log.Fatal(err)
	}

	db = database
	m.Run()

	err = database.Close()
	if err != nil {
		log.Printf("failed to close db connection: %s", err)
	}

	err = container.Terminate(ctx)
	if err != nil {
		log.Printf("failed to terminate the test container: %s", err)
	}
}
