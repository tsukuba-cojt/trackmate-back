package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type IExpenseCategoryRepository interface {
	FindAllExpenseCategory(userId string) (*[]models.ExpenseCategory, error)
	CreateExpenseCategory(expenseCategory models.ExpenseCategory) (*models.ExpenseCategory, error)
}

type ExpenseCategoryRepository struct {
	db *gorm.DB
}

func NewExpenseCategoryRepository(db *gorm.DB) IExpenseCategoryRepository {
	return &ExpenseCategoryRepository{db: db}
}

func (r *ExpenseCategoryRepository) FindAllExpenseCategory(userId string) (*[]models.ExpenseCategory, error) {
	var expenseCategories []models.ExpenseCategory
	result := r.db.Find(&expenseCategories, "user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &expenseCategories, nil
}

func (r *ExpenseCategoryRepository) CreateExpenseCategory(expenseCategory models.ExpenseCategory) (*models.ExpenseCategory, error) {
	result := r.db.Create(&expenseCategory)
	if result.Error != nil {
		return nil, result.Error
	}
	return &expenseCategory, nil
}
