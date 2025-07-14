package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"

	"github.com/google/uuid"
)

// インターフェースの定義
type ILoanPartnerService interface {
	FindAllLoanPartner(userId string) ([]models.LoanPerson, error)
	CreateLoanPartner(input dto.CreateLoanPersonInput) (*models.LoanPerson, error)
}

// サービスの定義
type LoanPartnerService struct {
	loanPartnerRepository repositories.ILoanPartnerRepository
}

// コンストラクタの定義
func NewLoanPartnerService(loanPartnerRepository repositories.ILoanPartnerRepository) ILoanPartnerService {
	return &LoanPartnerService{loanPartnerRepository: loanPartnerRepository}
}

// ユーザーごとの全ての借金の相手を取得する関数の定義
func (s *LoanPartnerService) FindAllLoanPartner(userId string) ([]models.LoanPerson, error) {
	return s.loanPartnerRepository.FindAllLoanPartner(userId)
}

// 借金の相手を作成する関数の定義
func (s *LoanPartnerService) CreateLoanPartner(input dto.CreateLoanPersonInput) (*models.LoanPerson, error) {
	newLoanPartnerID := uuid.New()
	newLoanPartner := models.LoanPerson{
		LoanPersonID:   newLoanPartnerID,
		UserID:         uuid.MustParse(input.UserID),
		LoanPersonName: input.CreateLoanPersonName,
	}

	return s.loanPartnerRepository.CreateLoanPartner(newLoanPartner)
}
