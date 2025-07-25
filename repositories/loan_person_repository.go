package repositories

import (
	"errors"
	"myapp/dto"
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type ILoanPersonRepository interface {
	FindAllLoanPerson(userId string) ([]dto.FindAllLoanPersonResponse, error)
	CreateLoanPerson(newLoanPerson models.LoanPerson) error
	DeleteLoanPerson(input dto.DeleteLoanPersonInput) error
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
func (r *LoanPersonRepository) FindAllLoanPerson(userId string) ([]dto.FindAllLoanPersonResponse, error) {
	var loanPersons []dto.FindAllLoanPersonResponse
	result := r.db.Table("loan_people").
		Select("loan_person_id as person_id, loan_person_name as person_name").
		Where("user_id = ? AND deleted_at IS NULL", userId).
		Scan(&loanPersons)
	if result.Error != nil {
		return nil, result.Error
	}
	return loanPersons, nil
}

// 借金の相手を作成する関数の定義
func (r *LoanPersonRepository) CreateLoanPerson(newLoanPerson models.LoanPerson) error {

	existingLoanPerson := models.LoanPerson{}
	result := r.db.Find(&existingLoanPerson, "loan_person_name = ? AND user_id = ?", newLoanPerson.LoanPersonName, newLoanPerson.UserID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		return errors.New("loan person already exists")
	}

	result = r.db.Create(&newLoanPerson)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *LoanPersonRepository) DeleteLoanPerson(input dto.DeleteLoanPersonInput) error {

	existingLoanPerson := models.LoanPerson{}
	result := r.db.Find(&existingLoanPerson, "loan_person_id = ? AND user_id = ? AND deleted_at IS NULL", input.PersonID, input.UserID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("loan person not found")
	}

	var count int64
	if err := r.db.Model(&models.Loan{}).
		Where("loan_person_id = ? AND deleted_at IS NULL", existingLoanPerson.LoanPersonID).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		// 参照されている場合はエラーを返す
		return errors.New("loan person is referenced by loans")
	}

	result = r.db.Delete(&models.LoanPerson{}, "loan_person_id = ? AND user_id = ? AND deleted_at IS NULL", input.PersonID, input.UserID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
