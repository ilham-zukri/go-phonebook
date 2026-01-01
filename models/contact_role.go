package models

import "github.com/google/uuid"

type ContactRole struct {
	ID       int       `json:"id" gorm:"primary_key" db:"id"`
	PublicID uuid.UUID `json:"public_id" db:"public_id"`
	Role     string    `json:"role" db:"role"`
}
