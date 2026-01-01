package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primary_key" db:"id"`
	PublicID  uuid.UUID      `json:"public_id" db:"public_id"`
	Password  string         `json:"password" db:"password" gorm:"column:password"`
	Name      string         `json:"name" db:"name"`
	EmpID     string         `json:"emp_id" db:"emp_id"`
	BranchID  int            `json:"branch_id" db:"branch_id"`
	RoleID    int            `json:"role_id" db:"role_id"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" db:"deleted_at"`
}

type UserResponse struct {
	PublicID  uuid.UUID `json:"public_id"`
	Name      string    `json:"name"`
	EmpID     string    `json:"emp_id"`
	BranchID  int       `json:"branch_id"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
