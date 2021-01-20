package models

type (
	//Inventory : struct representation of user inventory
	Inventory map[int]Item

	//Item : inventory item
	Item struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Timestamp   int64   `json:"timestamp"` //Item upload date
		Count       int     `json:"count"`     //Inventory count
		Status      int     `json:"status"`    //0:Pending, 1:Unlisted, 2:Unlisted, 3:Sold
	}

	//NewItem : new inventory item
	NewItem struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Timestamp   int64   `json:"timestamp"` //Item upload date
		Count       int     `json:"count"`     //Inventory count
	}
)
