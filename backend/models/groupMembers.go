package models

import (
	"errors"

	"gorm.io/gorm"
)

type Status string

const (
	CanCreateDebit    Status = "CAN_CREATE_DEBT"
	CannotCreateDebit Status = "CANNOT_CREATE_DEBT"
)

type GroupMembers struct {
	Id      int    `gorm:"primary_key"`
	UserId  int    `gorm:"not null"`
	GroupId int    `gorm:"not null"`
	Role    Status `gorm:"type:\"GroupMemberRole\";not null"`
	AuditableEntity
}

// Validate checks if the value is a valid enum
func (s Status) Validate() error {
	switch s {
	case CanCreateDebit, CannotCreateDebit:
		return nil
	default:
		return errors.New("invalid status value")
	}
}

// BeforeSave GORM hook to enforce enum validation
func (s *Status) BeforeSave(tx *gorm.DB) error {
	return s.Validate()
}