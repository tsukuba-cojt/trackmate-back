package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"
	"time"

	"github.com/google/uuid"
)

// インターフェースの定義
type IExpenseService interface {
	FindAllExpense(userId string) (*[]models.Expense, error)
	CreateExpense(input dto.CreateExpenseInput) (*models.Expense, error)
}

// サービスの定義
type ExpenseService struct {
	repository repositories.IExpenseRepository
}

// コンストラクタの定義
func NewExpenseService(repository repositories.IExpenseRepository) IExpenseService {
	return &ExpenseService{repository: repository}
}

// ユーザーごとの全ての支出を取得する関数の定義
func (s *ExpenseService) FindAllExpense(userId string) (*[]models.Expense, error) {
	return s.repository.FindAllExpense(userId)
}

// 支出を作成する関数の定義
func (s *ExpenseService) CreateExpense(input dto.CreateExpenseInput) (*models.Expense, error) {
	newExpenseID := uuid.New()
	expenseDate, err := time.Parse("2006-01-02", input.ExpenseDate)
	if err != nil {
		return nil, err
	}

	newExpense := models.Expense{
		ExpenseID:         newExpenseID,
		UserID:            uuid.MustParse(input.UserID),
		ExpenseCategoryID: uuid.MustParse(input.ExpenseCategoryID),
		ExpenseDate:       expenseDate,
		ExpenseAmount:     input.ExpenseAmount,
		Description:       input.Description,
	}

	return s.repository.CreateExpense(newExpense)
}
