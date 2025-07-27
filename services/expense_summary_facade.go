package services

import (
	"myapp/dto"
	"myapp/repositories"
	"time"
)

// ISummaryFacade はファサードのインターフェース
type ISummaryFacade interface {
	GetExpenseSummary(userId string) (*dto.ExpenseSummary, error)
	GetExpenseSummaryByDate(userId string, date string) (*[]dto.ExpenseSummaryByDate, error)
}

// SummaryFacade は複数のリポジトリを統合するファサード
type SummaryFacade struct {
	expenseRepository repositories.IExpenseRepository
	budgetRepository  repositories.IBudgetRepository
	loanRepository    repositories.ILoanRepository
}

// NewSummaryFacade はファサードのコンストラクタ
func NewSummaryFacade(
	expenseRepository repositories.IExpenseRepository,
	budgetRepository repositories.IBudgetRepository,
	loanRepository repositories.ILoanRepository,
) ISummaryFacade {
	return &SummaryFacade{
		expenseRepository: expenseRepository,
		budgetRepository:  budgetRepository,
		loanRepository:    loanRepository,
	}
}

// GetExpenseSummary は複数のリポジトリからデータを取得して統合する
func (f *SummaryFacade) GetExpenseSummary(userId string) (*dto.ExpenseSummary, error) {
	// 今月の期間を計算
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.UTC)
	firstOfNextMonth := firstOfMonth.AddDate(0, 1, 0)

	// 支出合計を取得
	expenseSum, err := f.expenseRepository.GetExpenseSum(userId, firstOfMonth, firstOfNextMonth)
	if err != nil {
		return nil, err
	}

	// 予算を取得
	budget, err := f.budgetRepository.FindBudgetByUserID(userId, firstOfMonth)
	if err != nil {
		return nil, err
	}

	// 負債を取得
	debtSum, err := f.loanRepository.GetDebtByUserID(userId)
	if err != nil {
		return nil, err
	}

	// 貸付を取得
	loanSum, err := f.loanRepository.GetLoanByUserID(userId)
	if err != nil {
		return nil, err
	}

	// サマリーを作成
	summary := &dto.ExpenseSummary{
		ExpensesSum:     expenseSum,
		Budget:          budget,
		RemainingBudget: budget - expenseSum,
		DebtSum:         debtSum,
		LoanSum:         loanSum,
	}

	return summary, nil
}

// GetExpenseSummaryByDate は特定日付の支出サマリーを取得
func (f *SummaryFacade) GetExpenseSummaryByDate(userId string, date string) (*[]dto.ExpenseSummaryByDate, error) {
	// 日付をパース
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	// その日の支出データを取得（期間を1日に限定）
	nextDay := parsedDate.AddDate(0, 0, 1)
	summary, err := f.expenseRepository.GetExpenseSummaryByDate(userId, parsedDate, nextDay)
	if err != nil {
		return nil, err
	}

	return summary, nil
}
