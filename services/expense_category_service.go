package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"

	"github.com/google/uuid"
)

// インターフェースの定義
type IExpenseCategoryService interface {
	GetExpenseCategorySummary(userId string) (*[]dto.ExpenseCategorySummaryResponse, error)
	CreateExpenseCategory(input dto.CreateExpenseCategoryInput) (*models.ExpenseCategory, error)
	DeleteExpenseCategory(userId string, expenseCategoryId string) error
}

// サービスの定義
type ExpenseCategoryService struct {
	repository repositories.IExpenseCategoryRepository
}

// コンストラクタの定義
func NewExpenseCategoryService(repository repositories.IExpenseCategoryRepository) IExpenseCategoryService {
	return &ExpenseCategoryService{repository: repository}
}

// ユーザーごとの全ての支出カテゴリを取得する関数の定義
func (s *ExpenseCategoryService) GetExpenseCategorySummary(userId string) (*[]dto.ExpenseCategorySummaryResponse, error) {
	return s.repository.GetExpenseCategorySummary(userId)
}

// 支出カテゴリを作成する関数の定義
func (s *ExpenseCategoryService) CreateExpenseCategory(input dto.CreateExpenseCategoryInput) (*models.ExpenseCategory, error) {
	newExpenseCategoryID := uuid.New()
	newExpenseCategory := models.ExpenseCategory{
		ExpenseCategoryID:   newExpenseCategoryID,
		UserID:              uuid.MustParse(input.UserID),
		ExpenseCategoryName: input.ExpenseCategoryName,
	}

	return s.repository.CreateExpenseCategory(newExpenseCategory)
}

func (s *ExpenseCategoryService) DeleteExpenseCategory(userId string, expenseCategoryId string) error {
	return s.repository.DeleteExpenseCategory(userId, expenseCategoryId)
}
