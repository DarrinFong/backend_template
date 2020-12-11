package services

import (
	"github.com/darrinfong/backend_template/dbinterface"
	"github.com/darrinfong/backend_template/models"
	"github.com/darrinfong/backend_template/tests"
)

var inventoryDB dbinterface.InventoryDB

//SetInventoryDB : 0-DB, 1-Mock
func SetInventoryDB(repoType int) {
	// Use DatabaseRepo or MockRepo depending on env.
	switch repoType {
	case 0:
		inventoryDB = dbinterface.SQLDB{}
	case 1:
		inventoryDB = tests.MockDB{}
	}
}

// GetInventory : Get seller inventory from DB and return as JSON
func GetInventory() map[int]models.Item {
	mockInventory := inventoryDB.GetInventory()
	mappedInventory := make(map[int]models.Item)
	for _, item := range mockInventory {
		mappedInventory[item.ID] = item
	}
	return mappedInventory
}

// GetItem : Get item by ID
func GetItem(itemID int) models.Item {
	return inventoryDB.GetItem(itemID)
}

// CreateItem : Create item
func CreateItem(newItem models.NewItem) models.Item {
	item := &models.Item{
		Name:        newItem.Name,
		Description: newItem.Description,
		Price:       newItem.Price,
		Timestamp:   newItem.Timestamp,
		Count:       newItem.Count,
		Status:      0,
	}
	return inventoryDB.CreateItem(item)
}
