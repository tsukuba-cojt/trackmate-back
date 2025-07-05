package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type IDebtPersonRepository interface {
	FindAllDebtPerson(userId string) ([]models.DebtPerson, error)
	CreateDebtPerson(newDebtPerson models.DebtPerson) (*models.DebtPerson, error)
}

// リポジトリの定義
type DebtPersonRepository struct {
	db *gorm.DB
}

// コンストラクタの定義
func NewDebtPersonRepository(db *gorm.DB) IDebtPersonRepository {
	return &DebtPersonRepository{db: db}
}

// ユーザーごとの全ての借金の相手を取得する関数の定義
func (r *DebtPersonRepository) FindAllDebtPerson(userId string) ([]models.DebtPerson, error) {
	var debtPersons []models.DebtPerson
	err := r.db.Where("user_id = ?", userId).Find(&debtPersons).Error
	return debtPersons, err
}

// 借金の相手を作成する関数の定義
func (r *DebtPersonRepository) CreateDebtPerson(newDebtPerson models.DebtPerson) (*models.DebtPerson, error) {
	result := r.db.Create(&newDebtPerson)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newDebtPerson, nil
}
