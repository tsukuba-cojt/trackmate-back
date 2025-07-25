package repositories

import (
	"errors"
	"myapp/models"
	"time"
	
	"gorm.io/gorm"
)

// インターフェースの定義
type IExpenseRepository interface {
	FindAllExpense(userId string) (*[]models.Expense, error)
	CreateExpense(expense models.Expense) (*models.Expense, error)
	FindExpensesByUserIDAndDateRange(userId string, startDate, endDate time.Time) ([]models.Expense, error)
	FindBudgetByUserID(userId string) (*models.Budget, error)
	FindDebtsByUserID(userId string) ([]models.Debt, error)
	FindLoansByUserID(userId string) ([]models.Loan, error)
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

// ユーザーごとの全ての支出を取得する関数の定義
func (r *ExpenseRepository) FindAllExpense(userId string) (*[]models.Expense, error) {
	var expense []models.Expense
	result := r.db.Find(&expense, "user_id = ?", userId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return &[]models.Expense{}, nil
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

// ユーザーIDと日付範囲で支出を取得する関数の定義
func (r *ExpenseRepository) FindExpensesByUserIDAndDateRange(userId string, startDate, endDate time.Time) ([]models.Expense, error) {
	var expenses []models.Expense
	result := r.db.
		Where("user_id = ?", userId).
		Where("expense_date >= ? AND expense_date < ?", startDate, endDate).
		Find(&expenses)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []models.Expense{}, nil
		}
		return nil, result.Error
	}
	return expenses, nil
}

// ユーザーの予算を取得する関数の定義
func (r *ExpenseRepository) FindBudgetByUserID(userId string) (*models.Budget, error) {
	var budget models.Budget
	result := r.db.Where("user_id = ?", userId).First(&budget)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &budget, nil
}

// ユーザーの負債を取得する関数の定義
func (r *ExpenseRepository) FindDebtsByUserID(userId string) ([]models.Debt, error) {
	var debts []models.Debt
	result := r.db.Where("user_id = ?", userId).Find(&debts)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []models.Debt{}, nil
		}
		return nil, result.Error
	}
	return debts, nil
}

// ユーザーの貸付を取得する関数の定義
func (r *ExpenseRepository) FindLoansByUserID(userId string) ([]models.Loan, error) {
	var loans []models.Loan
	result := r.db.Where("user_id = ?", userId).Find(&loans)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []models.Loan{}, nil // 見つからない場合は空のスライスを返す
		}
		return nil, result.Error
	}
	return loans, nil
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