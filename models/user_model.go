package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	Location  string             `json:"location,omitempty" validate:"required"`
	Title     string             `json:"title,omitempty" validate:"required"`
	DeletedAt time.Time          `json:"-" bson:"deleted_at,omitempty"`
}
