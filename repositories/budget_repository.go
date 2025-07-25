package repositories

import (
	"errors"
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type IBudgetRepository interface {
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
