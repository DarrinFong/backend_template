package inventory

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//DatabaseRepo : Implements InventoryRepo
type DatabaseRepo struct{}

//GetItem : Get item by id from DB
func (r DatabaseRepo) GetItem(id, userID int) *Item {
	item := &Item{ID: id}
	return item
}

//GetInventory : Get user inventory from DB
func (r DatabaseRepo) GetInventory(userID int) Inventory {
	inventory := make([]*Item, 10)
	for i := 0; i < 10; i++ {
		inventory[i] = r.GetItem(i, userID)
	}
	return inventory
}

//______________________________Mocks______________________________

//MockRepo : Used for unit tests and testing without DB.
type MockRepo struct{}

//GetItem : Return mock item.
func (r MockRepo) GetItem(id, userID int) *Item {
	// return pointer to mock item
	mockItem := &Item{
		ID:          id,
		SellerID:    rand.Intn(3),
		Name:        "JUNK" + fmt.Sprint(rand.Intn(10)),
		Description: "THIS IS MY JUNK.",
		Price:       math.Round(rand.Float64()*10000) / 100,
		Timestamp:   time.Now(),
		Status:      1,
	}
	return mockItem
}

//GetInventory : Return slice of mock items, count: number of items
func (r MockRepo) GetInventory(count int) Inventory {
	mockItems := make([]*Item, count)
	for i := 0; i < count; i++ {
		mockItems[i] = r.GetItem(i, count)
	}
	return mockItems
}
