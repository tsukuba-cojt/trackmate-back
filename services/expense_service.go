package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"
	"time"

	"github.com/google/uuid"
)

type IExpenseService interface {
	FindAllExpense(userId string) (*[]models.Expense, error)
	CreateExpense(input dto.CreateExpenseInput) (*models.Expense, error)
}

type ExpenseService struct {
	repository repositories.IExpenseRepository
}

func NewExpenseService(repository repositories.IExpenseRepository) IExpenseService {
	return &ExpenseService{repository: repository}
}

func (s *ExpenseService) FindAllExpense(userId string) (*[]models.Expense, error) {
	return s.repository.FindAllExpense(userId)
}

func (s *ExpenseService) CreateExpense(input dto.CreateExpenseInput) (*models.Expense, error) {
	newExpenseID := uuid.New()
	expenseDate, err := time.Parse("2006-01-02", input.ExpenseDate)
	if err != nil {
		return nil, err
	}

	newExpense := models.Expense{
		ExpenseID:         newExpenseID,
		UserID:            uuid.MustParse(input.UserID),
		ExpenseCategoryID: uuid.MustParse(input.ExpenseCategoryID),
		ExpenseDate:       expenseDate,
		ExpenseAmount:     input.ExpenseAmount,
	}

	return s.repository.CreateExpense(newExpense)
}
