package dbinterface

import (
	"github.com/darrinfong/backend_template/models"
)

//InventoryDB : managing interface for inventory
type InventoryDB interface {
	GetItem(int) *models.Item
	GetInventory(int) models.Inventory
	CreateItem(*models.Item)
	UpdateItem(int *models.Item)
}

//SQLDB : Implements InventoryDB
type SQLDB struct{}

//GetItem : Get item by id from DB
func (r SQLDB) GetItem(id int) *models.Item {
	item := &models.Item{ID: id}
	return item
}

//GetInventory : Get user inventory from DB
func (r SQLDB) GetInventory(userID int) models.Inventory {
	inventory := make([]*models.Item, 10)
	for i := 0; i < 10; i++ {
		inventory[i] = r.GetItem(i)
	}
	return inventory
}

func (r SQLDB) CreateItem(item *models.Item) {

}

func (r SQLDB) UpdateItem(int *models.Item) {

}