package repositories

import (
	"errors"
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type IExpenseRepository interface {
	FindAllExpense(userId string) (*[]models.Expense, error)
	CreateExpense(expense models.Expense) (*models.Expense, error)
}

// リポジトリの定義
type ExpenseRepository struct {
	db *gorm.DB
}

// コンストラクタの定義
func NewExpenseRepository(db *gorm.DB) IExpenseRepository {
	return &ExpenseRepository{db: db}
}

// ユーザーごとの全ての支出を取得する関数の定義
func (r *ExpenseRepository) FindAllExpense(userId string) (*[]models.Expense, error) {
	var expense []models.Expense
	result := r.db.Find(&expense, "user_id = ?", userId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("expense data not found")
		}
		return nil, result.Error
	}
	return &expense, nil
}

// 支出を作成する関数の定義
func (r *ExpenseRepository) CreateExpense(newExpense models.Expense) (*models.Expense, error) {
	result := r.db.Create(&newExpense)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newExpense, nil
}
