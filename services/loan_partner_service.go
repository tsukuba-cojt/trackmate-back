package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"

	"github.com/google/uuid"
)

// インターフェースの定義
type ILoanPartnerService interface {
	FindAllLoanPartner(userId string) ([]models.LoanPartner, error)
	CreateLoanPartner(input dto.CreateLoanPartnerInput) (*models.LoanPartner, error)
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
func (s *LoanPartnerService) FindAllLoanPartner(userId string) ([]models.LoanPartner, error) {
	return s.loanPartnerRepository.FindAllLoanPartner(userId)
}

// 借金の相手を作成する関数の定義
func (s *LoanPartnerService) CreateLoanPartner(input dto.CreateLoanPartnerInput) (*models.LoanPartner, error) {
	newLoanPartnerID := uuid.New()
	newLoanPartner := models.LoanPartner{
		LoanPartnerID:   newLoanPartnerID,
		UserID:          uuid.MustParse(input.UserID),
		LoanPartnerName: input.CreateLoanPartnerName,
	}

	return s.loanPartnerRepository.CreateLoanPartner(newLoanPartner)
}
