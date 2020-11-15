package pg

import (
	"context"
	"testing"
)

func TestItemListAssociation(t *testing.T) {
	ctx := context.Background()
	itemRepository := NewItemRepository(db)
	listRepository := NewListRepository(db)
	itemListRepository := NewItemListRepostory(db)

	createdItem, err := itemRepository.CreateItem(ctx, "desc")
	if err != nil {
		t.Error(err)
	}

	createdList, err := listRepository.CreateList(ctx, "title")
	if err != nil {
		t.Error(err)
	}

	err = itemListRepository.AddItemToList(ctx, createdItem.Id, createdList.Id)
	if err != nil {
		t.Error(err)
	}

	retrievedItems, err := itemListRepository.GetItemsByListId(ctx, createdList.Id)
	if err != nil {
		t.Error(err)
	}
	if retrievedItems[0].Id != createdItem.Id {
		t.Errorf("retrievedItems[0].Id != createdItem.Id (%s != %s)", retrievedItems[0].Id, createdItem.Id)
	}
	if retrievedItems[0].Description != createdItem.Description {
		t.Errorf("retrievedItems[0].Description != createdItem.Description (%s != %s)", retrievedItems[0].Description, createdItem.Description)
	}

	retrievedLists, err := itemListRepository.GetListsByItemId(ctx, createdItem.Id)
	if err != nil {
		t.Error(err)
	}
	if retrievedLists[0].Id != createdList.Id {
		t.Errorf("retrievedLists[0].Id != createdList.Id (%s != %s)", retrievedLists[0].Id, createdList.Id)
	}
	if retrievedLists[0].Title != createdList.Title {
		t.Errorf("retrievedLists[0].Title != createdList.Title (%s != %s)", retrievedLists[0].Title, createdList.Title)
	}
}
