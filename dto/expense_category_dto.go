package dto

// 支出カテゴリの作成の入力データの定義
type CreateExpenseCategoryInput struct {
	UserID              string `json:"user_id"`
	ExpenseCategoryName string `json:"expense_category_name" binding:"required"`
}
