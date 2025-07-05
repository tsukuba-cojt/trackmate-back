package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type IExpenseCategoryRepository interface {
	FindAllExpenseCategory(userId string) (*[]models.ExpenseCategory, error)
	CreateExpenseCategory(expenseCategory models.ExpenseCategory) (*models.ExpenseCategory, error)
	FindExpenseCategory(expenseCategoryID string) (string, error)
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
func (r *ExpenseCategoryRepository) FindExpenseCategory(expenseCategoryID string) (string, error) {
	expenseCategoryName := ""
	result := r.db.Find(&expenseCategoryName, "expense_category_id = ?", expenseCategoryID)
	if result.Error != nil {
		return "", result.Error
	}
	return expenseCategoryName, nil
}
