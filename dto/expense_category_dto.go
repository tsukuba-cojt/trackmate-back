package dto

type CreateExpenseCategoryInput struct {
	ExpenseCategoryName string `json:"expense_category_name" binding:"required"`
}
