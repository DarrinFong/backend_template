package tests

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/darrinfong/backend_template/models"
)

//MockDB : Used for unit tests and testing without DB.
type MockDB struct{}

//GetItem : Return mock item.
func (r MockDB) GetItem(id int) models.Item {
	// return pointer to mock item
	mockItem := models.Item{
		ID:          id,
		Name:        "STUFF" + fmt.Sprint(rand.Intn(10)),
		Description: "THIS IS MY STUFF.",
		Price:       math.Round(rand.Float64()*10000) / 100,
		Timestamp:   time.Now(),
		Count:       rand.Intn(10),
		Status:      1,
	}
	return mockItem
}

//GetInventory : Return slice of mock items, count: number of items
func (r MockDB) GetInventory() models.Inventory {
	mockInventory := make(map[int]models.Item)
	for i := 0; i < 10; i++ {
		mockInventory[i] = r.GetItem(i)
	}
	return mockInventory
}

func (r MockDB) CreateItem(item *models.Item) models.Item {
	item.ID = rand.Intn(9999999)
	return *item
}

func (r MockDB) UpdateItem(int *models.Item) {

}
