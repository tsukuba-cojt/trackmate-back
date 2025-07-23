package repositories

import (
	"errors"
	"myapp/models"
	"myapp/dto"

	"gorm.io/gorm"
)

// インターフェースの定義
type IExpenseRepository interface {
	FindAllExpense(userId string) (*[]models.Expense, error)
	CreateExpense(expense models.Expense) (*models.Expense, error)
	GetExpenseSummary(userId string) (*[]dto.GetExpenseSummary, error)
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

// ユーザーごとの支出のサマリーを取得する関数の定義
func (r *ExpenseRepository) GetExpenseSummary(userId string) (*[]dto.GetExpenseSummary, error) {
	var expenses []models.Expense
	result := r.db.Where("user_id = ?", userId).Find(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}

	var summary []dto.GetExpenseSummary
	for _, expense := range expenses {
		summary = append(summary, dto.GetExpenseSummary{
			ExpenseID:         expense.ExpenseID.String(),
			ExpenseCategoryID: expense.ExpenseCategoryID.String(),
			ExpenseDate:       expense.ExpenseDate.Format("2006-01-02"),
			ExpenseAmount:     expense.ExpenseAmount,
			Description:       expense.Description,
		})
	}

	return &summary, nil
}

// 支出を削除する関数の定義
func (r *ExpenseRepository) DeleteExpense(userId string, expenseId string) error {
	var expense models.Expense
	result := r.db.Where("user_id = ? AND expense_id = ?", userId, expenseId).First(&expense)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.New("expense not found")
		}
		return result.Error
	}

	result = r.db.Delete(&expense)
	if result.Error != nil {
		return result.Error
	}

	return nil
}