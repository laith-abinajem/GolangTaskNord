package model

type Branch struct {
	ID       uint   `gorm:"primaryKey,autoIncrement" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	TenantID uint   `gorm:"column:tenant_id" json:"tenant_id"`
}
