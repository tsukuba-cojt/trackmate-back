package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type IDebtPersonRepository interface {
	FindAllDebtPerson(userId string) ([]models.DebtPerson, error)
	CreateDebtPerson(newDebtPerson models.DebtPerson) (*models.DebtPerson, error)
}

type DebtPersonRepository struct {
	db *gorm.DB
}

func NewDebtPersonRepository(db *gorm.DB) IDebtPersonRepository {
	return &DebtPersonRepository{db: db}
}

func (r *DebtPersonRepository) FindAllDebtPerson(userId string) ([]models.DebtPerson, error) {
	var debtPersons []models.DebtPerson
	err := r.db.Where("user_id = ?", userId).Find(&debtPersons).Error
	return debtPersons, err
}

func (r *DebtPersonRepository) CreateDebtPerson(newDebtPerson models.DebtPerson) (*models.DebtPerson, error) {
	result := r.db.Create(&newDebtPerson)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newDebtPerson, nil
}
