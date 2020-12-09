package inventory

import (
	"encoding/json"
)

//InvRepo : managing interface for inventory
type InvRepo interface {
	GetItem(id, userID int) *Item
	GetInventory(int) Inventory
}

var InventoryRepo InvRepo

// ListUserInventory : Get user inventory from repo and return as JSON
func ListUserInventory(userID int) ([]byte, error) {
	mockInventory := InventoryRepo.GetInventory(10)
	return json.Marshal(mockInventory)
}
