package dto

type CreateExpenseInput struct {
	UserID            string `json:"user_id"`
	ExpenseCategoryID string `json:"expense_category_id" binding:"required"`
	ExpenseDate       string `json:"expense_date" binding:"required"`
	ExpenseAmount     int    `json:"expense_amount" binding:"required"`
	Description       string `json:"description"`
}
