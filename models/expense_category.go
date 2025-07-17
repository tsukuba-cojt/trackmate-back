package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 支出カテゴリのモデルの定義
type ExpenseCategory struct {
	ExpenseCategoryID   uuid.UUID `gorm:"type:char(36);primaryKey;unique"`
	UserID              uuid.UUID `gorm:"not null;foreignKey:UserID"`
	ExpenseCategoryName string    `gorm:"not null"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
