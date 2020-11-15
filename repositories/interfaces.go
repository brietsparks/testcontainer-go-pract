package repositories

import (
	"context"
	"pract-testcontainers/entities"
)

type ItemRepository interface {
	CreateItem(ctx context.Context, description string) (*entities.Item, error)
	GetItem(ctx context.Context, id string) (*entities.Item, error)
}

type ListRepository interface {
	CreateList(ctx context.Context, title string) (*entities.List, error)
	GetList(ctx context.Context, id string) (*entities.List, error)
}

type ItemListRepository interface {
	AddItemToList(ctx context.Context, itemId string, listId string) error
	GetItemsByListId(ctx context.Context, listId string) ([]*entities.Item, error)
	GetListsByItemId(ctx context.Context, itemId string) ([]*entities.List, error)
}
