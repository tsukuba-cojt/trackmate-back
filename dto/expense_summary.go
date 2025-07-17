package dto

type ExpenseSummaryResponse struct {
	ExpensesSum     int `json:"expenses_sum"`
	BudgetSum       int `json:"budget_sum"`
	RemainingBudget int `json:"remaining_budget"`
	DebtSum         int `json:"debt_sum"`
	LoanSum         int `json:"loan_sum"`
}
