package services

import (
	"encoding/json"

	dbinterface "github.com/darrinfong/backend_template/repository"
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

// ListUserInventory : Get user inventory from DB and return as JSON
func ListUserInventory(userID int) ([]byte, error) {
	mockInventory := inventoryDB.GetInventory(10)
	return json.Marshal(mockInventory)
}
