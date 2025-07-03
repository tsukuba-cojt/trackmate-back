package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"

	"github.com/google/uuid"
)

type IDebtPersonService interface {
	FindAllDebtPerson(userId string) ([]models.DebtPerson, error)
	CreateDebtPerson(input dto.CreateDebtPersonInput) (*models.DebtPerson, error)
}

type DebtPersonService struct {
	debtPersonRepository repositories.IDebtPersonRepository
}

func NewDebtPersonService(debtPersonRepository repositories.IDebtPersonRepository) IDebtPersonService {
	return &DebtPersonService{debtPersonRepository: debtPersonRepository}
}

func (s *DebtPersonService) FindAllDebtPerson(userId string) ([]models.DebtPerson, error) {
	return s.debtPersonRepository.FindAllDebtPerson(userId)
}

func (s *DebtPersonService) CreateDebtPerson(input dto.CreateDebtPersonInput) (*models.DebtPerson, error) {
	newDebtPersonID := uuid.New()
	newDebtPerson := models.DebtPerson{
		DebtPersonID:   newDebtPersonID,
		UserID:         uuid.MustParse(input.UserID),
		DebtPersonName: input.CreateDebtPersonName,
	}

	return s.debtPersonRepository.CreateDebtPerson(newDebtPerson)
}
