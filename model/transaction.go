package model

type Transaction struct {
	ID           uint    `gorm:"primaryKey,autoIncrement" json:"id"`
	TenantID     uint    `gorm:"column:tenant_id" json:"tenant_id"`
	BranchID     uint    `gorm:"column:branch_id" json:"branch_id"`
	ProductID    uint    `gorm:"column:product_id" json:"product_id"`
	Quantity     uint    `gorm:"column:quantity" json:"quantity"`
	PricePerUnit float64 `gorm:"column:price_per_unit" json:"price_per_unit"`
	TotalAmount  float64 `gorm:"column:total_amount" json:"total_amount"`
	CommonModel
}
