package models

import "github.com/google/uuid"

type Branch struct {
	ID       int       `json:"id" db:"id" gorm:"primary_key;auto_increment"`
	PublicID uuid.UUID `json:"public_id" db:"public_id"`
	RegionID string    `json:"region" db:"region_id"`
	Code     int       `json:"code" db:"code"`
	Name     string    `json:"name" db:"name"`
	Phone    string    `json:"phone" db:"phone"`
}
