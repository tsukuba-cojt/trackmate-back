package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type IExpenseCategoryRepository interface {
	FindAllExpenseCategory() (*[]models.ExpenseCategory, error)
}

type ExpenseCategoryRepository struct {
	db *gorm.DB
}

func NewExpenseCategoryRepository(db *gorm.DB) IExpenseCategoryRepository {
	return &ExpenseCategoryRepository{db: db}
}

func (r *ExpenseCategoryRepository) FindAllExpenseCategory() (*[]models.ExpenseCategory, error) {
	var expenseCategories []models.ExpenseCategory
	err := r.db.Find(&expenseCategories).Error
	return &expenseCategories, err
}
