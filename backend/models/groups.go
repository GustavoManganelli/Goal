package models

type Groups struct {
	Id            int             `gorm:"primaryKey"`
	Name          string          `gorm:"size:150;not null"`
	GroupExpenses []GroupExpenses `gorm:"foreignKey:GroupId;references:Id"`
	AuditableEntity
}