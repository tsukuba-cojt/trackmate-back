package services

import (
	"myapp/models"
	"myapp/repositories"
)

type IExpenseCategoryService interface {
	FindAllExpenseCategory() (*[]models.ExpenseCategory, error)
}

type ExpenseCategoryService struct {
	repository repositories.IExpenseCategoryRepository
}

func NewExpenseCategoryService(repository repositories.IExpenseCategoryRepository) IExpenseCategoryService {
	return &ExpenseCategoryService{repository: repository}
}

func (s *ExpenseCategoryService) FindAllExpenseCategory() (*[]models.ExpenseCategory, error) {
	return s.repository.FindAllExpenseCategory()
}
