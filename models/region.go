package models

type Region struct {
	ID      int    `json:"id" gorm:"primary_key" db:"id"`
	Region    string `json:"region" db:"region"`
}