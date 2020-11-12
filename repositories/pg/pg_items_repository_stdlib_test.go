package pg

import (
	"context"
	"testing"
)

func TestCreateItems(t *testing.T) {
	ctx := context.Background()
	r := NewItemsRepository(db)

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
		t.Errorf("created.Description (%s) != retrieved.Description (%s)", created.Description, retrieved.Description)
	}
}
