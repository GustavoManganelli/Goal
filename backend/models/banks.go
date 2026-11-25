package models

type Banks struct {
	Id       int        `gorm:"primary_key"`
	Name     string     `gorm:"size:150;not null"`
	Accounts []Accounts `gorm:"foreignKey:BankId"`
	AuditableEntity
}