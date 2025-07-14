package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type ILoanPartnerRepository interface {
	FindAllLoanPartner(userId string) ([]models.LoanPerson, error)
	CreateLoanPartner(newLoanPartner models.LoanPerson) (*models.LoanPerson, error)
}

// リポジトリの定義
type LoanPartnerRepository struct {
	db *gorm.DB
}

// コンストラクタの定義
func NewLoanPartnerRepository(db *gorm.DB) ILoanPartnerRepository {
	return &LoanPartnerRepository{db: db}
}

// ユーザーごとの全ての借金の相手を取得する関数の定義
func (r *LoanPartnerRepository) FindAllLoanPartner(userId string) ([]models.LoanPerson, error) {
	var loanPartners []models.LoanPerson
	err := r.db.Where("user_id = ?", userId).Find(&loanPartners).Error
	return loanPartners, err
}

// 借金の相手を作成する関数の定義
func (r *LoanPartnerRepository) CreateLoanPartner(newLoanPartner models.LoanPerson) (*models.LoanPerson, error) {
	result := r.db.Create(&newLoanPartner)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newLoanPartner, nil
}
