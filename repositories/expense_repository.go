package repositories

import (
	"errors"
	"myapp/dto"
	"myapp/models"
	"time"

	"gorm.io/gorm"
)

// インターフェースの定義
type IExpenseRepository interface {
	GetExpenseSum(userId string, startDate, endDate time.Time) (int, error)
	GetExpenseSummaryByDate(userId string, startDate, endDate time.Time) (*[]dto.ExpenseSummaryByDate, error)
	CreateExpense(expense models.Expense) (*models.Expense, error)
	DeleteExpense(userId string, expenseId string) error
}

// リポジトリの定義
type ExpenseRepository struct {
	db *gorm.DB
}

// コンストラクタの定義
func NewExpenseRepository(db *gorm.DB) IExpenseRepository {
	return &ExpenseRepository{db: db}
}

// ユーザーIDと日付範囲で支出の合計金額を取得する関数の定義
func (r *ExpenseRepository) GetExpenseSum(userId string, startDate, endDate time.Time) (int, error) {
	var sum int
	result := r.db.Table("expenses").
		Select("COALESCE(SUM(expense_amount), 0)").
		Where("user_id = ? AND expense_date BETWEEN ? AND ? AND deleted_at IS NULL", userId, startDate, endDate).
		Scan(&sum)
	if result.Error != nil {
		return 0, result.Error
	}
	return sum, nil
}

// ユーザーIDと日付範囲で支出のサマリーを取得する関数の定義
func (r *ExpenseRepository) GetExpenseSummaryByDate(userId string, startDate, endDate time.Time) (*[]dto.ExpenseSummaryByDate, error) {
	var summary []dto.ExpenseSummaryByDate
	result := r.db.Table("expenses").
		Select("expenses.expense_id as expense_id, expenses.expense_amount as expense_amount, expense_categories.expense_category_name as category_name").
		Joins("JOIN expense_categories ON expenses.expense_category_id = expense_categories.expense_category_id").
		Where("expenses.user_id = ? AND expenses.expense_date BETWEEN ? AND ? AND expenses.deleted_at IS NULL", userId, startDate, endDate).
		Scan(&summary)
	if result.Error != nil {
		return nil, result.Error
	}
	return &summary, nil
}

// 支出を作成する関数の定義
func (r *ExpenseRepository) CreateExpense(newExpense models.Expense) (*models.Expense, error) {
	result := r.db.Create(&newExpense)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newExpense, nil
}

// 支出を削除する関数の定義
func (r *ExpenseRepository) DeleteExpense(userId string, expenseId string) error {
	result := r.db.Where("user_id = ? AND expense_id = ?", userId, expenseId).Delete(&models.Expense{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("expense not found or not authorized to delete") // 削除対象が見つからない場合
	}
	return nil
}
