package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type Item struct {
	Id          string
	Description string
}

type ItemRepository struct {
	db *sql.DB
}

const createItem = `
insert into items(id, description) 
values ($1, $2)`

func (r *ItemRepository) CreateItem(ctx context.Context, description string) (*Item, error) {
	id := uuid.New()
	_, err := r.db.ExecContext(ctx, createItem, id, description)
	if err != nil {
		return nil, fmt.Errorf("item repository failed to create item: %s", err)
	}
	return &Item{Id: id.String(), Description: description}, nil
}

const getItem = `
select id, description 
from items
where id = $1`

func (r *ItemRepository) GetItem(ctx context.Context, id string) (*Item, error) {
	var item Item
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
