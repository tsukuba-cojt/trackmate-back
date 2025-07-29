package services

import (
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"
	"time"

	"github.com/google/uuid"
)

// インターフェースの定義
type IBudgetService interface {
	CreateBudget(input dto.CreateBudgetInput) error
}

// サービスの定義
type BudgetService struct {
	repository repositories.IBudgetRepository
}

// コンストラクタの定義
func NewBudgetService(repository repositories.IBudgetRepository) IBudgetService {
	return &BudgetService{repository: repository}
}

// 予算を作成する関数の定義
func (s *BudgetService) CreateBudget(input dto.CreateBudgetInput) error {
	newBudgetID := uuid.New()
	budgetDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return err
	}

	// 月の1日に変換し、JSTタイムゾーンで統一
	year, month, _ := budgetDate.Date()
	budgetDate = time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	newBudget := models.Budget{
		BudgetID: newBudgetID,
		UserID:   uuid.MustParse(input.UserID),
		Amount:   uint(input.Budget),
		Date:     budgetDate, // UTCで保存
	}
	return s.repository.CreateBudget(newBudget)
}
