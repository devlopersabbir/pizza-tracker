package models

var (
	OrderStatuses = []string{"Order Placed", "Preparing", "Baking", "Quality Check", "Ready"}

	PizzaTypes = []string{
		"Margherita",
		"Capricciosa",
		"Pepperoni",
		"Hawaiian",
		"Hawaiian",
		"Vegetarian",
		"Vegan",
		"Gluten Free",
		"Low Carb",
		"Medium Carb",
		"High Carb",
	}
	PizzaSizes = []string{
		"Small",
		"Medium",
		"Large",
	}
)

type OrderModel struct {
	DB *gorm.DB
}

type Order struct {
	ID string `gorm:"primaryKey;size:14" json:"id"`
	Status string `gorm:"not null" json:"status"`
	CustomerName string `gorm:"not null" json:"customerName`
	Phone string `gorm:"not null" json:"phone"`
	Address string `gorm:"not null" json:"address"`
	Items []OrderItem `gorm:"foreinKey:OrderID" json:"pizzas`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct {
	ID string `gorm:"primaryKey;size:14" json:"id"`
	OrderID string `gorm:"index;size:14;not null" json:"orderId"`
	Size string `gorm:"not null" json:"size"`
	Pizza string `gorm:"not null" json:"pizza"`
}
