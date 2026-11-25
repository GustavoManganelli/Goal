package models

import "github.com/jackc/pgx/v5/pgtype"

type AuditableEntity struct {
	CreatedBy string            `gorm:"size:150;not null"`
	CreatedAt pgtype.Timestamp  `gorm:"not null"`
	UpdatedBy *string           `gorm:"size:150"`
	UpdatedAt *pgtype.Timestamp `gorm:"null"`
}