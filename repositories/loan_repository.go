package repositories

import (
	"errors"
	"myapp/dto"
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type ILoanRepository interface {
	GetLoanSummary(userId string) (*[]dto.LoanSummaryResponse, error)
	CreateLoan(newLoan models.Loan) error
	DeleteLoan(userId string, personName string, isDebt bool) error
	FindLoan(userId string, userName string, isDebt bool) error
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
		Where("loans.user_id = ? AND loans.deleted_at IS NULL", userId).
		Group("loan_people.loan_person_name, loans.is_debt").
		Scan(&loanSummary)
	if result.Error != nil {
		return nil, result.Error
	}

	//fmt.Println(loanSummary)

	for i := range loanSummary {
		var loanHistory []dto.LoanHistoryResponse
		result := r.db.Table("loans").
			Select("loans.loan_date as date, loans.loan_amount as amount").
			Joins("JOIN loan_people ON loans.loan_person_id = loan_people.loan_person_id").
			Where("loan_people.loan_person_name = ? AND loans.is_debt = ? AND loans.deleted_at IS NULL", loanSummary[i].PersonName, loanSummary[i].IsDebt).
			Order("loans.loan_date ASC").
			Scan(&loanHistory)
		loanSummary[i].History = loanHistory
		if result.Error != nil {
			return nil, result.Error
		}
	}

	return &loanSummary, nil
}

// 借金を作成する関数の定義
func (r *LoanRepository) CreateLoan(newLoan models.Loan) error {
	result := r.db.Create(&newLoan)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *LoanRepository) DeleteLoan(userId string, personName string, isDebt bool) error {
	err := r.FindLoan(userId, personName, isDebt)
	if err != nil {
		return errors.New("loan not found")
	}

	// loan_person_idを取得
	var loanPersonID string
	err = r.db.Table("loan_people").
		Select("loan_person_id").
		Where("loan_person_name = ?", personName).
		Scan(&loanPersonID).Error
	if err != nil {
		return err
	}

	result := r.db.Where("user_id = ? AND loan_person_id = ? AND is_debt = ?", userId, loanPersonID, isDebt).
		Delete(&models.Loan{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *LoanRepository) FindLoan(userId string, userName string, isDebt bool) error {
	var loan models.Loan
	result := r.db.Table("loans").
		Select("loans.*").
		Joins("JOIN loan_people ON loans.loan_person_id = loan_people.loan_person_id").
		Where("loans.user_id = ? AND loan_people.loan_person_name = ? AND loans.is_debt = ? AND loans.deleted_at IS NULL", userId, userName, isDebt).
		First(&loan)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
