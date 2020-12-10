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
func (r MockDB) GetItem(id int) *models.Item {
	// return pointer to mock item
	mockItem := &models.Item{
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
func (r MockDB) GetInventory(count int) models.Inventory {
	mockItems := make([]*models.Item, count)
	for i := 0; i < count; i++ {
		mockItems[i] = r.GetItem(i)
	}
	return mockItems
}

func (r MockDB) CreateItem(item *models.Item) {

}

func (r MockDB) UpdateItem(int *models.Item) {

}
