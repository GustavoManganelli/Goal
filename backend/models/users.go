package models

import "github.com/jackc/pgx/v5/pgtype"

type Users struct {
	Id            int             `gorm:"primary_key"`
	FirstName     string          `gorm:"size:150;not null"`
	LastName      string          `gorm:"size:150;not null"`
	BirthDate     pgtype.Date     `gorm:"type:date;not null"`
	Email         string          `gorm:"size:255;not null"`
	PasswordHash  string          `gorm:"not null"`
	GroupExpenses []GroupExpenses `gorm:"foreignKey:UserId;references:Id"`
	UserDebts     []UserDebt      `gorm:"foreignKey:UserId;references:Id"`
	AuditableEntity
}
