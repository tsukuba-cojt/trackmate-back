package services

import (
	"myapp/models"
	"myapp/repositories"
)

type IExpenseService interface {
	FindAll() (*[]models.Expense, error)
}

type ExpenseService struct {
	repository repositories.IExpenseRepository
}

func NewExpenseService(repository repositories.IExpenseRepository) IExpenseService {
	return &ExpenseService{repository: repository}
}

func (s *ExpenseService) FindAll() (*[]models.Expense, error) {
	return s.repository.FindAll()
}
