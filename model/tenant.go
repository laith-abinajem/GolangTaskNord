package model

type Tenant struct {
	ID       uint     `gorm:"primaryKey,autoIncrement" json:"id"`
	Name     string   `gorm:"column:name" json:"name"`
	Branches []Branch `gorm:"foreignKey:TenantID" json:"branches"`
}
