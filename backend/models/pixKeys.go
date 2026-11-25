package models

import (
	"errors"

	"gorm.io/gorm"
)

type Type string

// Enum values
const (
	TypeCPF    Type = "CPF"
	TypeCNPJ   Type = "CNPJ"
	TypePhone  Type = "Phone"
	TypeEmail  Type = "Email"
	TypeRandom Type = "Random"
)

type PixKeys struct {
	Id        int    `gorm:"primary_key"`
	PixKey    string `gorm:"not null;unique"`
	AccountId int    `gorm:"not null"`
	KeyType   Type   `gorm:"type:\"PixKeyType\";not null"`
	AuditableEntity
}

// Validate checks if the value is a valid enum
func (s Type) Validate() error {
	switch s {
	case TypeCPF, TypeCNPJ, TypePhone, TypeEmail, TypeRandom:
		return nil
	default:
		return errors.New("invalid type value")
	}
}

// BeforeSave GORM hook to enforce enum validation
func (s *Type) BeforeSave(tx *gorm.DB) error {
	return s.Validate()
}