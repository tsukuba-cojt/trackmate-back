package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type ILoanRepository interface {
	FindAllLoan() (*[]models.Loan, error)
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

// ユーザーごとの全ての借金を取得する関数の定義
func (r *LoanRepository) FindAllLoan() (*[]models.Loan, error) {
	var loans []models.Loan
	err := r.db.Find(&loans).Error
	return &loans, err
}

// 借金を作成する関数の定義
func (r *LoanRepository) CreateLoan(newLoan models.Loan) (*models.Loan, error) {
	result := r.db.Create(&newLoan)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newLoan, nil
}
