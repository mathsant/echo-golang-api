package models

import "time"

type Transaction struct {
	Value           float64   `json:"value,omitempty" validate:"required"`
	UserId          string    `json:"user_id,omitempty"`
	Category        string    `json:"category,omitempty" validate:"required"`
	DateTransaction time.Time `json:"date_transaction"`
	CreatedAt       time.Time `json:"created_at"`
}
