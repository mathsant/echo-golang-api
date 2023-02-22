package models

import "time"

type Transaction struct {
	Value     float64   `json:"value,omitempty" validate:"required"`
	UserId    string    `json:"user_id,omitempty" validate:"required"`
	Category  string    `json:"category,omitempty" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
