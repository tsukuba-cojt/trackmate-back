package dto

type CreateExpenseCategoryInput struct {
	UserID              string `json:"user_id"`
	ExpenseCategoryName string `json:"expense_category_name" binding:"required"`
}
