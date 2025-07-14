package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Budget struct {
	BudgetID  uuid.UUID      `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID      `gorm:"not null, foreignKey:UserID"`
	Amount    uint           `gorm:"not null"`
	Date      time.Time      `gorm:"not null"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
