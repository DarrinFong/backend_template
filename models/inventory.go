package models

import (
	"time"
)

type (
	//Inventory : struct representation of user inventory
	Inventory []*Item

	//Item : struct representation of an inventory item
	Item struct {
		ID          int       `json:"id"`
		SellerID    int       `json:"sellerId"`
		Name        string    `json:"name"`
		Description string    `json:"desc"`
		Price       float64   `json:"price"`
		Timestamp   time.Time `json:"timestamp"` //Item upload date
		Status      int       `json:"status"`    //0:Pending, 1:Unlisted, 2:Unlisted, 3:Sold
	}
)
