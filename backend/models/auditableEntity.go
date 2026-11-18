package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type AuditableEntity struct {
	createdBy string
	createdAt bson.DateTime
	updatedBy *string
	updatedAt *bson.DateTime
}
