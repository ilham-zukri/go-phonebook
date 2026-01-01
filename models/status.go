package models

import "github.com/google/uuid"

type Status struct {
	ID       int       `json:"id" gorm:"primary_key" db:"id"`
	PublicID uuid.UUID `json:"public_id" db:"public_id"`
	Status   string    `json:"status" db:"status"`
}
