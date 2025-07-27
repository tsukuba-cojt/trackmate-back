package repositories

import (
	"errors"
	"myapp/models"
	"time"

	"gorm.io/gorm"
)

// インターフェースの定義
type IBudgetRepository interface {
	FindBudgetByUserID(userId string, date time.Time) (int, error)
	CreateBudget(input models.Budget) error
}

// リポジトリの定義
type BudgetRepository struct {
	db *gorm.DB
}

// コンストラクタの定義
func NewBudgetRepository(db *gorm.DB) IBudgetRepository {
	return &BudgetRepository{db: db}
}

// ユーザーIDに紐づく予算を取得する関数の定義
func (r *BudgetRepository) FindBudgetByUserID(userId string, date time.Time) (int, error) {
	var budget int
	result := r.db.Table("budgets").
		Select("COALESCE(amount, 0)").
		Where("user_id = ? AND date = ?", userId, date).
		Scan(&budget)
	if result.Error != nil {
		return 0, result.Error
	}
	return budget, nil
}

// 予算を作成する関数の定義
func (r *BudgetRepository) CreateBudget(input models.Budget) error {
	existingBudget := models.Budget{}
	result := r.db.Find(&existingBudget, "date = ? AND user_id = ?", input.Date, input.UserID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		return errors.New("budget already exists")
	}
	result = r.db.Create(&input)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
