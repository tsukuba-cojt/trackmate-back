package repositories

import (
	"errors"
	"myapp/dto"
	"myapp/models"
	"gorm.io/gorm"
)

func (r *ExpenseRepository) ExpenseSummary(userId string) (*dto.ExpenseSummary, error) {
    var expenses []models.Expense
    result := r.db.Where("user_id = ?", userId).Find(&expenses)

    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return &dto.ExpenseSummary{
                ExpensesSum:     0,
                Budget:          0,
                RemainingBudget: 0,
                DebtSum:         0,
                LoanSum:         0,
            }, nil
        }
        return nil, result.Error
    }

    totalExpenses := 0
    for _, expense := range expenses {
        totalExpenses += expense.Amount
    }

	

    // ここに予算、負債、貸付などの情報を取得・計算するロジックを追加してください
    // 例: ユーザーの予算情報をデータベースから取得
    // var userBudget models.Budget
    // r.db.Where("user_id = ?", userId).First(&userBudget)
    // budget := userBudget.Amount

    // 仮のデータ（実際はDBや他の情報から取得・計算してください）

    summary := dto.ExpenseSummary{
        ExpensesSum:     totalExpenses,
        Budget:          budget,
        RemainingBudget: budget - totalExpenses,
        DebtSum:         debtSum,
        LoanSum:         loanSum,
    }

    return &summary, nil
}