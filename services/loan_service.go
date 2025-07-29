package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"
	"time"

	"github.com/google/uuid"
)

// インターフェースの定義
type ILoanService interface {
	GetLoanSummary(userId string) (*[]dto.LoanSummaryResponse, error)
	CreateLoan(input dto.CreateLoanInput) error
	DeleteLoan(userId string, personName string, isDebt *bool) error
}

// サービスの定義
type LoanService struct {
	repository repositories.ILoanRepository
}

// コンストラクタの定義
func NewLoanService(repository repositories.ILoanRepository) ILoanService {
	return &LoanService{repository: repository}
}

// ユーザーごとの全ての借金を取得する関数の定義
func (s *LoanService) GetLoanSummary(userId string) (*[]dto.LoanSummaryResponse, error) {
	return s.repository.GetLoanSummary(userId)
}

// 借金を作成する関数の定義
func (s *LoanService) CreateLoan(input dto.CreateLoanInput) error {
	newDebtID := uuid.New()
	loanDate, err := time.ParseInLocation("2006-01-02", input.LoanDate, time.Local)
	if err != nil {
		return err
	}
	// 日付のみの場合は、時刻を00:00:00 JSTに設定
	loanDate = time.Date(loanDate.Year(), loanDate.Month(), loanDate.Day(), 0, 0, 0, 0, time.Local)

	newLoan := models.Loan{
		LoanID:       newDebtID,
		UserID:       uuid.MustParse(input.UserID),
		LoanPersonID: uuid.MustParse(input.LoanPersonID),
		IsDebt:       *input.IsDebt,
		LoanDate:     loanDate,
		LoanAmount:   input.LoanAmount,
	}

	return s.repository.CreateLoan(newLoan)
}

// 借金を削除する関数の定義
func (s *LoanService) DeleteLoan(userId string, personName string, isDebt *bool) error {
	return s.repository.DeleteLoan(userId, personName, *isDebt)
}
