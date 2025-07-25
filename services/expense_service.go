package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"myapp/dto"
	"myapp/models"
	"myapp/repositories"
	"gorm.io/gorm"
)

// インターフェースの定義
type IExpenseService interface {
	FindAllExpense(userId string) ([]models.Expense, error)
	CreateExpense(input dto.CreateExpenseInput) (*models.Expense, error)
	ExpenseSummary(userId string) (*dto.ExpenseSummary, error)
	DeleteExpense(userId string, expenseId string) error
}

// サービスの定義
type ExpenseService struct {
	ExpenseRepository repositories.IExpenseRepository
}

// コンストラクタの定義
func NewExpenseService(expenseRepository repositories.IExpenseRepository) IExpenseService {
	return &ExpenseService{ExpenseRepository: expenseRepository}
}

// FindAllExpense はサービス層でリポジトリに委譲するメソッド
func (s *ExpenseService) FindAllExpense(userId string) ([]models.Expense, error) {
	expenses, err := s.ExpenseRepository.FindAllExpense(userId)
	if err != nil {
		return nil, err
	}
	if expenses == nil {
		return []models.Expense{}, nil
	}
	return *expenses, nil
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

	// ExpenseDate (string) を time.Time に変換
	parsedDate, err := time.Parse("2006-01-02", input.ExpenseDate)
	if err != nil {
		return nil, errors.New("invalid expense date format")
	}

	newExpense := models.Expense{
		UserID:            userID,
		ExpenseCategoryID: expenseCategoryID,
		ExpenseDate:       parsedDate,
		ExpenseAmount:     input.ExpenseAmount,
		Description:       input.Description,
	}

	createdExpense, err := s.ExpenseRepository.CreateExpense(newExpense)
	if err != nil {
		return nil, err
	}
	return createdExpense, nil
}


// ExpenseSummary はサービス層で集計ロジックを実行し、リポジトリからデータを取得する
func (s *ExpenseService) ExpenseSummary(userId string) (*dto.ExpenseSummary, error) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, now.Location())
	firstOfNextMonth := firstOfMonth.AddDate(0, 1, 0)

	// リポジトリから今月の支出データを取得
	expenses, err := s.ExpenseRepository.FindExpensesByUserIDAndDateRange(userId, firstOfMonth, firstOfNextMonth)
	if err != nil {
		return nil, err
	}

	totalExpenses := 0
	for _, expense := range expenses {
		totalExpenses += expense.ExpenseAmount
	}

	// リポジトリから予算、負債、貸付の合計情報を取得
	budgetModel, err := s.ExpenseRepository.FindBudgetByUserID(userId)
	budget := 0
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else if budgetModel != nil {
		budget = int(budgetModel.Amount)
	}

	debts, err := s.ExpenseRepository.FindDebtsByUserID(userId)
	debtSum := 0
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		// モデルに金額フィールドがないため、ここでは合計を計算できません。
		// 負債の合計を実際に算出するには、models.Debtに金額を表すフィールドが必要です。
		// 例: debtSum += int(debt.Amount) または debtSum += int(debt.DebtAmount)
		// 現在はモデルにフィールドがないため、加算は行わず debtSum は 0 のままです。
		for range debts { // ループは回るが、金額を加算するフィールドがない
			// debtSum += <ここに金額フィールドを記述>
		}
	}

	loans, err := s.ExpenseRepository.FindLoansByUserID(userId)
	loanSum := 0
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		// モデルに金額フィールドがないため、ここでは合計を計算できません。
		// 貸付の合計を実際に算出するには、models.Loanに金額を表すフィールドが必要です。
		// 例: loanSum += int(loan.Amount) または loanSum += int(loan.LoanAmount)
		// 現在はモデルにフィールドがないため、加算は行わず loanSum は 0 のままです。
		for range loans { // ループは回るが、金額を加算するフィールドがない
			// loanSum += <ここに金額フィールドを記述>
		}
	}

	summary := dto.ExpenseSummary{
		ExpensesSum:     totalExpenses,
		Budget:          budget,
		RemainingBudget: budget - totalExpenses,
		DebtSum:         debtSum,
		LoanSum:         loanSum,
	}

	return &summary, nil
}

// DeleteExpense はサービス層でリポジトリに委譲するメソッド
func (s *ExpenseService) DeleteExpense(userId string, expenseId string) error {
	err := s.ExpenseRepository.DeleteExpense(userId, expenseId)
	if err != nil {
		return err
	}
	return nil
}