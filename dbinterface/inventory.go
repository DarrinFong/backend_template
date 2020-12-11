package dbinterface

import (
	"github.com/darrinfong/backend_template/models"
)

//InventoryDB : managing interface for inventory
type InventoryDB interface {
	GetItem(int) models.Item
	GetInventory() models.Inventory
	CreateItem(*models.Item) models.Item
	UpdateItem(int *models.Item)
}

//SQLDB : Implements InventoryDB
type SQLDB struct{}

//GetItem : Get item by id from DB
func (r SQLDB) GetItem(id int) models.Item {
	item := models.Item{ID: id}
	return item
}

//GetInventory : Get user inventory from DB
func (r SQLDB) GetInventory() models.Inventory {
	inventory := make(map[int]models.Item)
	for i := 0; i < 10; i++ {
		inventory[i] = r.GetItem(i)
	}
	return inventory
}

func (r SQLDB) CreateItem(item *models.Item) models.Item {
	return models.Item{}
}

func (r SQLDB) UpdateItem(int *models.Item) {

}
