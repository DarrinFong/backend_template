package services

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/darrinfong/backend_template/dbinterface"
	"github.com/darrinfong/backend_template/models"
	"github.com/stretchr/testify/assert"
)

var i = models.Item{
	ID:          1,
	Name:        "thing1",
	Description: "Test thing1",
	Price:       11.11,
	Timestamp:   time.Now(),
	Count:       1,
	Status:      0,
}

var inv = models.Inventory{
	1: models.Item{
		ID:          1,
		Name:        "thing1",
		Description: "Test thing1",
		Price:       11.11,
		Timestamp:   time.Now(),
		Count:       1,
		Status:      1,
	},
	2: models.Item{
		ID:          2,
		Name:        "thing2",
		Description: "Test thing2",
		Price:       22.22,
		Timestamp:   time.Now(),
		Count:       2,
		Status:      2,
	},
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Could not create DB mock")
		log.Fatalf(err.Error())
	}
	return db, mock
}

func TestGetItem(t *testing.T) {
	db, mock := NewMock()
	repo := &dbinterface.SQLDB{DB: db}
	InventoryRepo = repo

	defer repo.DB.Close()

	expectedQuery := "SELECT id, name, descriptions, price, timestamp, count, status FROM inventory WHERE id = ?"

	rows := sqlmock.
		NewRows([]string{"id", "name", "descriptions", "price", "timestamp", "count", "status"}).
		AddRow(i.ID, i.Name, i.Description, i.Price, i.Timestamp, i.Count, i.Status)

	mock.ExpectQuery(expectedQuery).WithArgs(i.ID).WillReturnRows(rows)

	item, err := GetItem(i.ID)
	assert.NoError(t, err)
	assert.Equal(t, i, item)
}

func TestGetInventory(t *testing.T) {
	db, mock := NewMock()
	repo := &dbinterface.SQLDB{DB: db}
	InventoryRepo = repo

	defer repo.DB.Close()

	expectedQuery := "SELECT id, name, descriptions, price, timestamp, count, status FROM inventory "
	expectedRows := sqlmock.
		NewRows([]string{"id", "name", "descriptions", "price", "timestamp", "count", "status"}).
		AddRow(inv[1].ID, inv[1].Name, inv[1].Description, inv[1].Price, inv[1].Timestamp, inv[1].Count, inv[1].Status).
		AddRow(inv[2].ID, inv[2].Name, inv[2].Description, inv[2].Price, inv[2].Timestamp, inv[2].Count, inv[2].Status)

	mock.ExpectQuery(expectedQuery).WillReturnRows(expectedRows)

	inventory, err := GetInventory()
	assert.NoError(t, err)
	assert.Equal(t, inv, inventory)
}
