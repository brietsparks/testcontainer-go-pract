package pg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"pract-testcontainers/entities"
	"pract-testcontainers/repositories"
)

type ListRepository struct {
	db *sql.DB
}

func NewListRepository(db *sql.DB) repositories.ListRepository {
    return &ListRepository{db: db}
}

const createList = `
insert into lists(id, title)
values($1, $2)`

func (r *ListRepository) CreateList(ctx context.Context, title string) (*entities.List, error) {
	id := uuid.New()
	_, err := r.db.ExecContext(ctx, createList, id, title)
	if err != nil {
		return nil, fmt.Errorf("list repository failed to create list: %s", err)
	}
	return &entities.List{Id: id.String(), Title: title}, nil
}

const getList = `
select id, title
from lists
where id = $1`

func (r *ListRepository) GetList(ctx context.Context, id string) (*entities.List, error) {
	var list entities.List
	err := r.db.
		QueryRowContext(ctx, getList, id).
		Scan(&list.Id, &list.Title)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("list repository encountered an error while trying to get list: %s", err)
	}

	return &list, nil
}
