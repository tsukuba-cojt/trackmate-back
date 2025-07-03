package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Expense struct {
	ExpenseID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID            uuid.UUID `gorm:"not null;foreignKey:UserID"`
	ExpenseCategoryID uuid.UUID `gorm:"not null;foreignKey:ExpenseCategoryID"`
	ExpenseDate       time.Time
	ExpenseAmount     int `gorm:"not null"`
	Description       string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
