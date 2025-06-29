package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Debt struct {
	DebtID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID        uuid.UUID `gorm:"not null;foreignKey:UserID"`
	DebtPartnerID uuid.UUID `gorm:"not null;foreignKey:DebtPersonID"`
	IsBorrow      bool      `gorm:"not null"`
	DebtDate      time.Time
	DebtAmount    int `json:"amount"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
