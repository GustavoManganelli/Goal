package models

import "github.com/shopspring/decimal"

type UserDebt struct {
	Id             int             `gorm:"primaryKey"`
	Amount         decimal.Decimal `gorm:"type:numeric(10,2);not null"`
	UserId         int             `gorm:"not null"`
	GroupExpenseId int             `gorm:"not null"`
}