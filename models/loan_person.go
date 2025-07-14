package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 借金の相手のモデルの定義
type LoanPartner struct {
	LoanPartnerID   uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID          uuid.UUID `gorm:"not null;foreignKey:UserID"`
	LoanPartnerName string    `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
