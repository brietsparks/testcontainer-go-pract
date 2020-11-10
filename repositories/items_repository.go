package repositories

import (
	"pract-testcontainers/entities"
)

type ItemsRepository interface {
	CreateItem() (*entities.Item, error)
}
