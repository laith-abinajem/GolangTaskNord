package model

type Product struct {
	ID    uint    `gorm:"primaryKey,autoIncrement" json:"id"`
	Name  string  `gorm:"column:name" json:"name"`
	Price float64 `gorm:"column:price" json:"price"`
}
type TopProduct struct {
	ProductID     uint    `json:"product_id" gorm:"column:product_id"`
	Name          string  `json:"name" gorm:"column:name"`
	Price         float64 `json:"price" gorm:"column:price"`
	TotalQuantity uint    `json:"total_quantity" gorm:"column:total_quantity"`
}
