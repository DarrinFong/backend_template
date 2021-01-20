package services

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/darrinfong/backend_template/dbinterface"
	"github.com/darrinfong/backend_template/models"
	"github.com/stretchr/testify/assert"
)

var currentTimeUnix = time.Now().Unix()

var i = models.Item{
	ID:          0,
	Name:        "thing1",
	Description: "Test thing1",
	Price:       11.11,
	Timestamp:   currentTimeUnix,
	Count:       1,
	Status:      0,
}

var n = models.NewItem{
	Name:        "thing1",
	Description: "Test thing1",
	Price:       11.11,
	Timestamp:   currentTimeUnix,
	Count:       1,
}

var inv = models.Inventory{
	0: models.Item{
		ID:          0,
		Name:        "thing1",
		Description: "Test thing1",
		Price:       11.11,
		Timestamp:   currentTimeUnix,
		Count:       1,
		Status:      1,
	},
	1: models.Item{
		ID:          1,
		Name:        "thing2",
		Description: "Test thing2",
		Price:       22.22,
		Timestamp:   currentTimeUnix + 1,
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
	t.Parallel()
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
	t.Parallel()
	db, mock := NewMock()
	repo := &dbinterface.SQLDB{DB: db}
	InventoryRepo = repo

	defer repo.DB.Close()

	expectedQuery := "SELECT id, name, descriptions, price, timestamp, count, status FROM inventory "
	expectedRows := sqlmock.
		NewRows([]string{"id", "name", "descriptions", "price", "timestamp", "count", "status"}).
		AddRow(inv[0].ID, inv[0].Name, inv[0].Description, inv[0].Price, inv[0].Timestamp, inv[0].Count, inv[0].Status).
		AddRow(inv[1].ID, inv[1].Name, inv[1].Description, inv[1].Price, inv[1].Timestamp, inv[1].Count, inv[1].Status)

	mock.ExpectQuery(expectedQuery).WillReturnRows(expectedRows)

	inventory, err := GetInventory()
	assert.NoError(t, err)
	assert.Equal(t, inv, inventory)
}

func TestCreateItem(t *testing.T) {
	db, mock := NewMock()
	repo := &dbinterface.SQLDB{DB: db}
	InventoryRepo = repo

	defer repo.DB.Close()

	expectedQuery := `INSERT INTO inventory (name, descriptions, price, timestamp, count, status)
	VALUES (%q, %q, '%f', '%d', %q, %q)
	SELECT SCOPE_IDENTITY()`
	expectedQuery = fmt.Sprintf(expectedQuery, n.Name, n.Description, n.Price, n.Timestamp, n.Count, 0)
	expectedQuery = regexp.QuoteMeta(expectedQuery)
	expectedRows := sqlmock.
		NewRows([]string{"id", "name", "descriptions", "price", "timestamp", "count", "status"}).
		AddRow(i.ID, i.Name, i.Description, i.Price, i.Timestamp, i.Count, i.Status)

	mock.ExpectQuery(expectedQuery).WillReturnRows(expectedRows)

	item, err := CreateItem(n)
	assert.NoError(t, err)
	assert.Equal(t, i, item)
}
