package pg

import (
	"context"
	"database/sql"
	"fmt"
	"pract-testcontainers/entities"
	"pract-testcontainers/repositories"
)

type ItemListRepository struct {
	db    *sql.DB
}

func NewItemListRepostory(db *sql.DB) repositories.ItemListRepository {
    return &ItemListRepository{db}
}

const addItemToList = `
insert into items_lists(item_id, list_id) 
values ($1, $2)`

func (r *ItemListRepository) AddItemToList(ctx context.Context, itemId string, listId string) error {
	_, err := r.db.ExecContext(ctx, addItemToList, itemId, listId)
	if err != nil {
		return fmt.Errorf("item-list repository failed to add item to list: %s", err)
	}

	return nil
}

const getItemsByListId = `
select id, description
from items
inner join items_lists on items.id = items_lists.item_id
where items_lists.list_id = $1`

func (r *ItemListRepository) GetItemsByListId(ctx context.Context, listId string) ([]*entities.Item, error) {
	rows, err := r.db.QueryContext(ctx, getItemsByListId, listId)
	if err != nil {
		return nil, fmt.Errorf("item-list repository failed to retrieve items by listId: %s", err)
	}
	defer rows.Close()

	items := make([]*entities.Item, 0)
	for rows.Next() {
		item := entities.Item{}
		rows.Scan(&item.Id, &item.Description)
		items = append(items, &item)
	}

	return items, nil
}

const getListsByItemId = `
select id, title
from lists
inner join items_lists on lists.id = items_lists.list_id
where items_lists.item_id = $1`

func (r *ItemListRepository) GetListsByItemId(ctx context.Context, itemId string) ([]*entities.List, error) {
	rows, err := r.db.QueryContext(ctx, getListsByItemId, itemId)
	if err != nil {
		return nil, fmt.Errorf("item-list repository failed to retrieve lists by itemId: %s", err)
	}
	defer rows.Close()

	lists := make([]*entities.List, 0)
	for rows.Next() {
		list := entities.List{}
		rows.Scan(&list.Id, &list.Title)
		lists = append(lists, &list)
	}

	return lists, nil
}
