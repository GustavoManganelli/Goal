package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type Payments struct {
	Id              int              `gorm:"primaryKey"`
	Amount          decimal.Decimal  `gorm:"type:numeric(10,2);not null"`
	TransactionDate pgtype.Timestamp `gorm:"type:timestamp;not null"`
	UserDebtId      int              `gorm:"not null"`
	Proof           string           `gorm:"type:text;not null"`
}