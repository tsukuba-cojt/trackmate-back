package repositories

import (
	"errors"
	"myapp/models"

	"gorm.io/gorm"
)

type IExpenseRepository interface {
	FindAllExpense() (*[]models.Expense, error)
	CreateExpense(expense models.Expense) (*models.Expense, error)
}

type ExpenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) IExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) FindAllExpense() (*[]models.Expense, error) {
	var expense []models.Expense
	result := r.db.Find(&expense)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("expense data not found")
		}
		return nil, result.Error
	}
	return &expense, nil
}

func (r *ExpenseRepository) CreateExpense(newExpense models.Expense) (*models.Expense, error) {
	result := r.db.Create(&newExpense)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newExpense, nil
}
