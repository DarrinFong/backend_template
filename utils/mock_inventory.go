package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/darrinfong/backend_template/models"
)

//HMockDB : Used for unit tests and testing without DB.
type HMockDB struct{}

/* Deprecated hardcoded mocks */
func (r *HMockDB) GetItem(id int) (models.Item, error) {
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
	return mockItem, nil
}

func (r *HMockDB) GetInventory() (models.Inventory, error) {
	mockInventory := make(map[int]models.Item)
	for i := 0; i < 10; i++ {
		mockInventory[i], _ = r.GetItem(i)
	}
	return mockInventory, nil
}

func (r *HMockDB) CreateItem(item models.Item) (models.Item, error) {
	item.ID = rand.Intn(9999999)
	return item, nil
}

func (r *HMockDB) UpdateItem(id int, item models.Item) (models.Item, error) {
	item.ID = rand.Intn(9999999)
	return item, nil
}
