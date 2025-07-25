package services

import (
	"errors"
	"time"
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
	ExpenseRepository repositories.IExpenseRepository // リポジトリのインターフェース
}

// コンストラクタの定義
func NewExpenseService(expenseRepository repositories.IExpenseRepository) IExpenseService { // 戻り値をインターフェース型に
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
    newExpense := models.Expense{
        UserID:            input.UserID,
        ExpenseCategoryID: input.ExpenseCategoryID,
        ExpenseDate:       time.Time{},
        Amount:            input.ExpenseAmount,
        Description:       input.Description,
    }

    // ここで input.ExpenseDate (string) を time.Time に変換するロジックを追加
    parsedDate, err := time.Parse("2006-01-02", input.ExpenseDate) // "YYYY-MM-DD"形式と仮定
    if err != nil {
        return nil, errors.New("invalid expense date format")
    }
    newExpense.ExpenseDate = parsedDate


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
		totalExpenses += expense.Amount
	}

	// リポジトリから予算、負債、貸付の合計情報を取得
	budgetModel, err := s.ExpenseRepository.FindBudgetByUserID(userId)
	budget := 0
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) { // レコードが見つからないエラー以外は返す
			return nil, err
		}
	} else if budgetModel != nil {
		budget = budgetModel.Amount
	}

	debts, err := s.ExpenseRepository.FindDebtsByUserID(userId)
	debtSum := 0
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		for _, debt := range debts {
			debtSum += debt.Amount
		}
	}

	loans, err := s.ExpenseRepository.FindLoansByUserID(userId)
	loanSum := 0
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		for _, loan := range loans {
			loanSum += loan.Amount
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