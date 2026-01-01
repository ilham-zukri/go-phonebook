package models

import "github.com/google/uuid"

type Region struct {
	ID       int       `json:"id" gorm:"primary_key" db:"id"`
	PublicID uuid.UUID `json:"public_id" db:"public_id"`
	Region   string    `json:"region" db:"region"`
}
