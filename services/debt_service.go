package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"
	"time"

	"github.com/google/uuid"
)

type IDebtService interface {
	FindAllDebt() (*[]models.Debt, error)
	CreateDebt(input dto.CreateDebtInput) (*models.Debt, error)
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

func (s *DebtService) CreateDebt(input dto.CreateDebtInput) (*models.Debt, error) {
	newDebtID := uuid.New()
	debtDate, err := time.Parse("2006-01-02", input.DebtDate)
	if err != nil {
		return nil, err
	}

	newDebt := models.Debt{
		DebtID:       newDebtID,
		UserID:       uuid.MustParse(input.UserID),
		DebtPersonID: uuid.MustParse(input.DebtPersonID),
		IsBorrow:     input.IsBorrow,
		DebtDate:     debtDate,
		DebtAmount:   input.DebtAmount,
	}

	return s.repository.CreateDebt(newDebt)
}
