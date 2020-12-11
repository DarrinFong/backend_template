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

// GetSellerInventory : Get seller inventory from DB and return as JSON
func GetSellerInventory(sellerID int) map[int]models.Item {
	mockInventory := inventoryDB.GetInventory(10)
	sellerInventory := make(map[int]models.Item)
	for _, item := range mockInventory {
		if item.SellerID == sellerID {
			sellerInventory[item.ID] = item
		}
	}
	return sellerInventory
}

// GetItem : Get item by ID
func GetItem(itemID int) models.Item {
	return inventoryDB.GetItem(itemID)
}
