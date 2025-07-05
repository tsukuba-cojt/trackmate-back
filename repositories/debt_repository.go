package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type IDebtRepository interface {
	FindAllDebt() (*[]models.Debt, error)
	CreateDebt(newDebt models.Debt) (*models.Debt, error)
}

// リポジトリの定義
type DebtRepository struct {
	db *gorm.DB
}

// コンストラクタの定義
func NewDebtRepository(db *gorm.DB) IDebtRepository {
	return &DebtRepository{db: db}
}

// ユーザーごとの全ての借金を取得する関数の定義
func (r *DebtRepository) FindAllDebt() (*[]models.Debt, error) {
	var debts []models.Debt
	err := r.db.Find(&debts).Error
	return &debts, err
}

// 借金を作成する関数の定義
func (r *DebtRepository) CreateDebt(newDebt models.Debt) (*models.Debt, error) {
	result := r.db.Create(&newDebt)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newDebt, nil
}
