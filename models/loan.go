package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 借金のモデルの定義
type Loan struct {
	LoanID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID        uuid.UUID `gorm:"not null;foreignKey:UserID"`
	LoanPartnerID uuid.UUID `gorm:"not null;foreignKey:LoanPartnerID"`
	IsDebt        bool      `gorm:"not null"`
	LoanDate      time.Time
	LoanAmount    int `json:"amount"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
