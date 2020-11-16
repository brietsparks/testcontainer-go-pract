package main

import (
	"context"
	"testing"
)

func TestExample(t *testing.T)  {
	ctx := context.Background()

	// container and database
	container, db, err := CreateTestContainer(ctx, "testdb")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer container.Terminate(ctx)

	// migration
	mig, err := NewPgMigrator(db)
	if err != nil {
		t.Fatal(err)
	}

	err = mig.Up()
	if err != nil {
		t.Fatal(err)
	}

	// test
	r :=  &ItemRepository{db}
	created, err := r.CreateItem(ctx, "desc")
	if err != nil {
		t.Errorf("failed to create item: %s", err)
	}
	retrieved, err := r.GetItem(ctx, created.Id)
	if err != nil {
		t.Errorf("failed to retrieve item: %s", err)
	}
	if created.Id != retrieved.Id {
		t.Errorf("created.Id (%s) != retrieved.Id (%s)", created.Id, retrieved.Id)
	}
	if created.Description != retrieved.Description {
		t.Errorf("created.Description != retrieved.Description (%s != %s)", created.Description, retrieved.Description)
	}
}
