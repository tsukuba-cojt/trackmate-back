package services

import (
	"errors"
	"fmt"
	"time"

	"myapp/dto"
	"myapp/models"
	"myapp/repositories"

	"github.com/google/uuid"
)

// インターフェースの定義
type IExpenseService interface {
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

// CreateExpense はサービス層でリポジトリに委譲するメソッド
func (s *ExpenseService) CreateExpense(input dto.CreateExpenseInput) (*models.Expense, error) {
	// string UserID を uuid.UUID に変換
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, errors.New("invalid UserID format")
	}

	// string ExpenseCategoryID を uuid.UUID に変換
	expenseCategoryID, err := uuid.Parse(input.ExpenseCategoryID)
	if err != nil {
		return nil, errors.New("invalid ExpenseCategoryID format")
	}

	// ExpenseDate (string) を time.Time に変換（JSTで統一）
	parsedDate, err := time.ParseInLocation("2006-01-02", input.ExpenseDate, time.Local)
	if err != nil {
		return nil, errors.New("invalid expense date format")
	}
	// 日付のみの場合は、時刻を00:00:00 JSTに設定
	parsedDate = time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, time.Local)

	expenseId := uuid.New()

	newExpense := models.Expense{
		ExpenseID:         expenseId,
		UserID:            userID,
		ExpenseCategoryID: expenseCategoryID,
		ExpenseDate:       parsedDate,
		ExpenseAmount:     input.ExpenseAmount,
		Description:       input.Description,
	}

	fmt.Println(expenseId)

	createdExpense, err := s.repository.CreateExpense(newExpense)
	if err != nil {
		return nil, err
	}
	return createdExpense, nil
}

// DeleteExpense はサービス層でリポジトリに委譲するメソッド
func (s *ExpenseService) DeleteExpense(userId string, expenseId string) error {
	err := s.repository.DeleteExpense(userId, expenseId)
	if err != nil {
		return err
	}
	return nil
}
