package dbinterface

import (
	"database/sql"
	"fmt"

	"github.com/darrinfong/backend_template/models"
)

//SQLDB : Implements InventoryDB
type SQLDB struct {
	DB *sql.DB
}

//GetItem : Get item by id from DB
func (r *SQLDB) GetItem(id int) (models.Item, error) {
	query := "SELECT id, name, descriptions, price, timestamp, count, status FROM inventory WHERE id = ?"
	result := r.DB.QueryRow(query, id)

	i := models.Item{}
	err := result.Scan(&i.ID, &i.Name, &i.Description, &i.Price, &i.Timestamp, &i.Count, &i.Status)
	return i, err
}

//GetInventory : Get user inventory from DB
func (r *SQLDB) GetInventory() (models.Inventory, error) {
	query := "SELECT id, name, descriptions, price, timestamp, count, status FROM inventory"
	resultRows, err := r.DB.Query(query)
	if err != nil {
		return models.Inventory{}, err
	}
	inv, err := mapRowsToInventory(resultRows)
	return inv, err
}

//CreateItem :
func (r *SQLDB) CreateItem(item models.Item) (models.Item, error) {
	query := `INSERT INTO inventory (name, descriptions, price, timestamp, count, status)
	VALUES (%q, %q, '%f', '%d', %q, %q)
	SELECT SCOPE_IDENTITY()`
	query = fmt.Sprintf(query, item.Name, item.Description, item.Price, item.Timestamp, item.Count, item.Status)

	result := r.DB.QueryRow(query)
	err := result.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Timestamp, &item.Count, &item.Status)
	return item, err
}

//UpdateItem :
func (r *SQLDB) UpdateItem(id int, item models.Item) (models.Item, error) {
	query := `UPDATE inventory
	SET name = $item.Name, description = $item.Description, price = $item.Price, timestamp = $item.Timestamp, count = $item.Count, status = $item.Status
	WHERE id = $id`
	println(query)
	return models.Item{ID: 1}, nil
}

func mapRowsToInventory(rows *sql.Rows) (models.Inventory, error) {
	inv := models.Inventory{}
	for rows.Next() {
		i := models.Item{}
		err := rows.Scan(&i.ID, &i.Name, &i.Description, &i.Price, &i.Timestamp, &i.Count, &i.Status)
		if err != nil {
			return models.Inventory{}, err
		}
		inv[i.ID] = i
	}
	return inv, nil
}
