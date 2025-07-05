package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ユーザーのモデルの定義
type User struct {
	UserID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
