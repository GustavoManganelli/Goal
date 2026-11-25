package models

type Accounts struct {
	Id     int `gorm:"primary_key"`
	UserId int `gorm:"not null"`
	BankId int `gorm:"not null"`
	AuditableEntity
}
