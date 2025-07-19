package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 借金の相手のモデルの定義
type LoanPerson struct {
	LoanPersonID   uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID         uuid.UUID `gorm:"not null;foreignKey:UserID"`
	LoanPersonName string    `gorm:"not null;unique"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
