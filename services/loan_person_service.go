package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"

	"github.com/google/uuid"
)

// インターフェースの定義
type ILoanPersonService interface {
	FindAllLoanPerson(userId string) ([]models.LoanPerson, error)
	CreateLoanPerson(input dto.CreateLoanPersonInput) (*models.LoanPerson, error)
}

// サービスの定義
type LoanPersonService struct {
	loanPersonRepository repositories.ILoanPersonRepository
}

// コンストラクタの定義
func NewLoanPersonService(loanPersonRepository repositories.ILoanPersonRepository) ILoanPersonService {
	return &LoanPersonService{loanPersonRepository: loanPersonRepository}
}

// ユーザーごとの全ての借金の相手を取得する関数の定義
func (s *LoanPersonService) FindAllLoanPerson(userId string) ([]models.LoanPerson, error) {
	return s.loanPersonRepository.FindAllLoanPerson(userId)
}

// 借金の相手を作成する関数の定義
func (s *LoanPersonService) CreateLoanPerson(input dto.CreateLoanPersonInput) (*models.LoanPerson, error) {
	newLoanPersonID := uuid.New()
	newLoanPerson := models.LoanPerson{
		LoanPersonID:   newLoanPersonID,
		UserID:         uuid.MustParse(input.UserID),
		LoanPersonName: input.CreateLoanPersonName,
	}

	return s.loanPersonRepository.CreateLoanPerson(newLoanPerson)
}
