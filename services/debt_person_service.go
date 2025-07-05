package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"

	"github.com/google/uuid"
)

// インターフェースの定義
type IDebtPersonService interface {
	FindAllDebtPerson(userId string) ([]models.DebtPerson, error)
	CreateDebtPerson(input dto.CreateDebtPersonInput) (*models.DebtPerson, error)
}

// サービスの定義
type DebtPersonService struct {
	debtPersonRepository repositories.IDebtPersonRepository
}

// コンストラクタの定義
func NewDebtPersonService(debtPersonRepository repositories.IDebtPersonRepository) IDebtPersonService {
	return &DebtPersonService{debtPersonRepository: debtPersonRepository}
}

// ユーザーごとの全ての借金の相手を取得する関数の定義
func (s *DebtPersonService) FindAllDebtPerson(userId string) ([]models.DebtPerson, error) {
	return s.debtPersonRepository.FindAllDebtPerson(userId)
}

// 借金の相手を作成する関数の定義
func (s *DebtPersonService) CreateDebtPerson(input dto.CreateDebtPersonInput) (*models.DebtPerson, error) {
	newDebtPersonID := uuid.New()
	newDebtPerson := models.DebtPerson{
		DebtPersonID:   newDebtPersonID,
		UserID:         uuid.MustParse(input.UserID),
		DebtPersonName: input.CreateDebtPersonName,
	}

	return s.debtPersonRepository.CreateDebtPerson(newDebtPerson)
}
