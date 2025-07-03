package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type IDebtRepository interface {
	FindAllDebt() (*[]models.Debt, error)
	CreateDebt(newDebt models.Debt) (*models.Debt, error)
}

type DebtRepository struct {
	db *gorm.DB
}

func NewDebtRepository(db *gorm.DB) IDebtRepository {
	return &DebtRepository{db: db}
}

func (r *DebtRepository) FindAllDebt() (*[]models.Debt, error) {
	var debts []models.Debt
	err := r.db.Find(&debts).Error
	return &debts, err
}

func (r *DebtRepository) CreateDebt(newDebt models.Debt) (*models.Debt, error) {
	result := r.db.Create(&newDebt)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newDebt, nil
}
