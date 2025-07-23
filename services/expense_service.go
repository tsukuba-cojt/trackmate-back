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
	GetExpenseSummary(userId string) (*[]dto.GetExpenseCategorySummary, error)
	CreateExpense(input dto.CreateExpenseInput) (*models.Expense, error)
	DeleteExpense(userId string, expenseId string) error
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
func (s *ExpenseService) GetExpenseSummary(userId string) (*[]dto.GetExpenseSummary, error) {
	return s.repository.GetExpenseSummary(userId)
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

// 支出を削除する関数の定義
func (s *ExpenseService) DeleteExpense(userId string, expenseId string) error {
	expenseUUID, err := uuid.Parse(expenseId)
	if err != nil {
		return err
	}

	return s.repository.DeleteExpense(userId, expenseUUID)
}