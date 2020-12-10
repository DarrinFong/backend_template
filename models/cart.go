package models

type (
	CartItem struct {
		ItemCount int
		Item      *Item
	}

	Cart struct {
		UserID     int
		Items      []*CartItem
		TotalPrice float32
	}
)
