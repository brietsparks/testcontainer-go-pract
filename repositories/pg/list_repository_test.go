package pg

import (
	"context"
	"testing"
)

func TestCreateAndGetList(t *testing.T) {
	ctx := context.Background()
	r := NewListRepository(db)

	created, err := r.CreateList(ctx, "desc")
	if err != nil {
		t.Errorf("failed to create list: %s", err)
	}

	retrieved, err := r.GetList(ctx, created.Id)
	if err != nil {
		t.Errorf("failed to retrieve list: %s", err)
	}
	if created.Id != retrieved.Id {
		t.Errorf("created.Id (%s) != retrieved.Id (%s)", created.Id, retrieved.Id)
	}

	if created.Title != retrieved.Title {
		t.Errorf("created.Description (%s) != retrieved.Description (%s)", created.Title, retrieved.Title)
	}
}
