package services

import (
	"myapp/models"
	"myapp/repositories"
)

type IDebtService interface {
	FindAllDebt() (*[]models.Debt, error)
}

type DebtService struct {
	repository repositories.IDebtRepository
}

func NewDebtService(repository repositories.IDebtRepository) IDebtService {
	return &DebtService{repository: repository}
}

func (s *DebtService) FindAllDebt() (*[]models.Debt, error) {
	return s.repository.FindAllDebt()
}
