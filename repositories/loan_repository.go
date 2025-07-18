package repositories

import (
	"myapp/dto"
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type ILoanRepository interface {
	GetLoanSummary(userId string) (*[]dto.LoanSummaryResponse, error)
	CreateLoan(newLoan models.Loan) (*models.Loan, error)
}

// リポジトリの定義
type LoanRepository struct {
	db *gorm.DB
}

// コンストラクタの定義
func NewLoanRepository(db *gorm.DB) ILoanRepository {
	return &LoanRepository{db: db}
}

// ユーザーごとの借金の合計金額を取得する関数の定義
func (r *LoanRepository) GetLoanSummary(userId string) (*[]dto.LoanSummaryResponse, error) {
	var loanSummary []dto.LoanSummaryResponse
	result := r.db.Table("loans").
		Select("loan_people.loan_person_name as person_name, loans.is_debt as is_debt, SUM(loan_amount) as sum_amount").
		Joins("JOIN loan_people ON loans.loan_person_id = loan_people.loan_person_id").
		Where("user_id = ?", userId).
		Group("loan_people.loan_person_name, loans.is_debt").
		Scan(&loanSummary)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := range loanSummary {
		var loanHistory []dto.LoanHistoryResponse
		result := r.db.Table("loans").
			Select("loan_date, loan_amount").
			Where("loan_person_id = ? AND is_debt = ?", loanSummary[i].PersonName, loanSummary[i].IsDebt).
			Scan(&loanHistory)
		loanSummary[i].History = loanHistory
		if result.Error != nil {
			return nil, result.Error
		}
	}

	return &loanSummary, nil
}

// 借金を作成する関数の定義
func (r *LoanRepository) CreateLoan(newLoan models.Loan) (*models.Loan, error) {
	result := r.db.Create(&newLoan)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newLoan, nil
}
