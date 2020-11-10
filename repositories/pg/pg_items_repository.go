package pg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"pract-testcontainers/entities"
)

type ItemsRepository struct {
	db *sql.DB
}

func NewItemsRepository(db *sql.DB) *ItemsRepository {
	return &ItemsRepository{db: db}
}

const createItem = `
	insert into items(id, description) 
	values ($1, $2)`

func (r *ItemsRepository) CreateItem(ctx context.Context, description string) (*entities.Item, error) {
	id := uuid.New()
	_, err := r.db.ExecContext(ctx, createItem, id, description)
	if err != nil {
		return nil, fmt.Errorf("item repository failed to create item: %s", err)
	}
	return &entities.Item{Id: id.String(), Description: description}, nil
}

const getItem = `
	select id, description 
	from items
	where id = $1`

func (r *ItemsRepository) GetItem(ctx context.Context, id string) (*entities.Item, error) {
	var item entities.Item
    err := r.db.
    	QueryRowContext(ctx, getItem, id).
    	Scan(&item.Id, &item.Description)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("item repository encountered an error while trying to get item: %s", err)
	}

	return &item, nil
}
