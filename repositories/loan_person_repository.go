package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type ILoanPersonRepository interface {
	FindAllLoanPerson(userId string) ([]models.LoanPerson, error)
	CreateLoanPerson(newLoanPerson models.LoanPerson) (*models.LoanPerson, error)
}

// リポジトリの定義
type LoanPersonRepository struct {
	db *gorm.DB
}

// コンストラクタの定義
func NewLoanPersonRepository(db *gorm.DB) ILoanPersonRepository {
	return &LoanPersonRepository{db: db}
}

// ユーザーごとの全ての借金の相手を取得する関数の定義
func (r *LoanPersonRepository) FindAllLoanPerson(userId string) ([]models.LoanPerson, error) {
	var loanPersons []models.LoanPerson
	err := r.db.Where("user_id = ?", userId).Find(&loanPersons).Error
	return loanPersons, err
}

// 借金の相手を作成する関数の定義
func (r *LoanPersonRepository) CreateLoanPerson(newLoanPerson models.LoanPerson) (*models.LoanPerson, error) {
	result := r.db.Create(&newLoanPerson)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newLoanPerson, nil
}
