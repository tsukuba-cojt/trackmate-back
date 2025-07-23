package dto

// 支出の作成の入力データ
type CreateExpenseInput struct {
    UserID          string `json:"user_id"`
    ExpenseDate     string `json:"expense_date" binding:"required"`
    ExpenseAmount   int    `json:"expense_amount" binding:"required"`
    ExpenseCategoryID string `json:"category_id" binding:"required"`
    Description     string `json:"description"`
}

// 支出のサマリーのレスポンス
type ExpenseSummary struct {
    ExpensesSum     int `json:"expenses_sum"`
    Budget          int `json:"budget"`
    RemainingBudget int `json:"remaining_budget"`
    DebtSum         int `json:"debt_sum"`
    LoanSum         int `json:"loan_sum"`
}

// 日ごとの支出のレスポンス
type ExpenseByDate struct {
    ExpenseId       int    `json:"expense_id"`
    ExpenseAmount   int    `json:"expense_amount"`
    ExpenseCategory string `json:"expense_category"`
}

// 支出削除のレスポンス
type DeleteExpenseResponse struct {
    ExpenseID string `json:"expense_id"`
}