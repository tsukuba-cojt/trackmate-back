package dto

// 支出カテゴリの作成の入力データの定義
type CreateExpenseCategoryInput struct {
	UserID              string `json:"user_id"`
	ExpenseCategoryName string `json:"expense_category_name" binding:"required"`
}

type ExpenseCategorySummaryResponse struct {
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	Sum          int    `json:"sum"`
}

type DeleteExpenseCategoryInput struct {
	ExpenseCategoryID string `json:"category_id" binding:"required"`
}
