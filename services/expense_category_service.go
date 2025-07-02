package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"

	"github.com/google/uuid"
)

type IExpenseCategoryService interface {
	FindAllExpenseCategory(userId string) (*[]models.ExpenseCategory, error)
	CreateExpenseCategory(input dto.CreateExpenseCategoryInput) (*models.ExpenseCategory, error)
}

type ExpenseCategoryService struct {
	repository repositories.IExpenseCategoryRepository
}

func NewExpenseCategoryService(repository repositories.IExpenseCategoryRepository) IExpenseCategoryService {
	return &ExpenseCategoryService{repository: repository}
}

func (s *ExpenseCategoryService) FindAllExpenseCategory(userId string) (*[]models.ExpenseCategory, error) {
	return s.repository.FindAllExpenseCategory(userId)
}

func (s *ExpenseCategoryService) CreateExpenseCategory(input dto.CreateExpenseCategoryInput) (*models.ExpenseCategory, error) {
	newExpenseCategoryID := uuid.New()
	newExpenseCategory := models.ExpenseCategory{
		ExpenseCategoryID:   newExpenseCategoryID,
		UserID:              uuid.MustParse(input.UserID),
		ExpenseCategoryName: input.ExpenseCategoryName,
	}

	return s.repository.CreateExpenseCategory(newExpenseCategory)
}
