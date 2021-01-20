package services

import (
	"github.com/darrinfong/backend_template/models"
)

//InventoryInterface : managing interface for inventory
type InventoryInterface interface {
	GetItem(int) (models.Item, error)
	GetInventory() (models.Inventory, error)
	CreateItem(models.Item) (models.Item, error)
	UpdateItem(int, models.Item) (models.Item, error)
}

//InventoryRepo :
var InventoryRepo InventoryInterface

// GetInventory : Get seller inventory from DB and return as JSON
func GetInventory() (models.Inventory, error) {
	inv, err := InventoryRepo.GetInventory()
	return inv, err
}

// GetItem : Get item by ID
func GetItem(itemID int) (models.Item, error) {
	return InventoryRepo.GetItem(itemID)
}

// CreateItem : Create item
func CreateItem(newItem models.NewItem) (models.Item, error) {
	item := models.Item{
		Name:        newItem.Name,
		Description: newItem.Description,
		Price:       newItem.Price,
		Timestamp:   newItem.Timestamp,
		Count:       newItem.Count,
		Status:      0,
	}
	return InventoryRepo.CreateItem(item)
}

// UpdateItem : Modify existing item
func UpdateItem(itemID int, modifiedItem models.Item) (models.Item, error) {
	return InventoryRepo.UpdateItem(itemID, modifiedItem)
}
