package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type IDebtRepository interface {
	FindAllDebt() (*[]models.Debt, error)
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
