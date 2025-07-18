package repositories

import (
	"errors"
	"myapp/dto"
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type IExpenseCategoryRepository interface {
	FindAllExpenseCategory(userId string) (*[]models.ExpenseCategory, error)
	CreateExpenseCategory(expenseCategory models.ExpenseCategory) (*models.ExpenseCategory, error)
	FindExpenseCategory(userId string, expenseCategoryId string) (*models.ExpenseCategory, error)
	GetExpenseCategorySummary(userId string) (*[]dto.ExpenseCategorySummaryResponse, error)
	DeleteExpenseCategory(userId string, expenseCategoryId string) error
}

// リポジトリの定義
type ExpenseCategoryRepository struct {
	db *gorm.DB
}

// コンストラクタの定義
func NewExpenseCategoryRepository(db *gorm.DB) IExpenseCategoryRepository {
	return &ExpenseCategoryRepository{db: db}
}

// ユーザーごとの全ての支出カテゴリを取得する関数の定義
func (r *ExpenseCategoryRepository) FindAllExpenseCategory(userId string) (*[]models.ExpenseCategory, error) {
	var expenseCategories []models.ExpenseCategory
	result := r.db.Find(&expenseCategories, "user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &expenseCategories, nil
}

// 支出カテゴリを作成する関数の定義
func (r *ExpenseCategoryRepository) CreateExpenseCategory(expenseCategory models.ExpenseCategory) (*models.ExpenseCategory, error) {
	result := r.db.Create(&expenseCategory)
	if result.Error != nil {
		return nil, result.Error
	}
	return &expenseCategory, nil
}

// 支出カテゴリを取得する関数の定義
func (r *ExpenseCategoryRepository) FindExpenseCategory(userId string, expenseCategoryId string) (*models.ExpenseCategory, error) {
	expenseCategory := models.ExpenseCategory{}
	result := r.db.Find(&expenseCategory, "expense_category_id = ? AND user_id = ?", expenseCategoryId, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &expenseCategory, nil
}

// 支出カテゴリの合計金額を取得する関数の定義
func (r *ExpenseCategoryRepository) GetExpenseCategorySummary(userId string) (*[]dto.ExpenseCategorySummaryResponse, error) {
	var expenseCategorySummary []dto.ExpenseCategorySummaryResponse
	result := r.db.Table("expenses").
		Select("expenses.expense_category_id as category_id, expense_categories.expense_category_name as category_name, SUM(expenses.expense_amount) as sum").
		Joins("JOIN expense_categories ON expenses.expense_category_id = expense_categories.expense_category_id").
		Where("expenses.user_id = ? AND expenses.deleted_at IS NULL", userId).
		Group("expenses.expense_category_id, expense_categories.expense_category_name").
		Scan(&expenseCategorySummary)
	if result.Error != nil {
		return nil, result.Error
	}
	return &expenseCategorySummary, nil
}

func (r *ExpenseCategoryRepository) DeleteExpenseCategory(userId string, expenseCategoryId string) error {
	_, err := r.FindExpenseCategory(userId, expenseCategoryId)
	if err != nil {
		return errors.New("expense category not found")
	}

	var count int64
	if err := r.db.Model(&models.Expense{}).
		Where("expense_category_id = ?", expenseCategoryId).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		// 参照されている場合はエラーを返す
		return errors.New("category is referenced by expenses")
	}

	result := r.db.Where("expense_category_id = ? AND user_id = ?", expenseCategoryId, userId).Delete(&models.ExpenseCategory{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
