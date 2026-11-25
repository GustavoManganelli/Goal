package models

import "github.com/shopspring/decimal"

type GroupExpenses struct {
	Id          int             `gorm:"primaryKey"`
	Subject     string          `gorm:"size:150;not null"`
	TotalAmount decimal.Decimal `gorm:"type:numeric(10,2);not null"`
	GroupId     int             `gorm:"not null"`
	UserId      int             `gorm:"not null"`
	AuditableEntity
}