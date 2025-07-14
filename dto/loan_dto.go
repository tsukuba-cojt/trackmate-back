package dto

// 借金の作成の入力データの定義
type CreateLoanInput struct {
	UserID       string `json:"user_id"`
	LoanPersonID string `json:"loan_person_id" binding:"required"`
	IsDebt       bool   `json:"is_debt" binding:"required"`
	LoanDate     string `json:"loan_date" binding:"required"`
	LoanAmount   int    `json:"amount" binding:"required"`
	Description  string `json:"description" binding:"required"`
}
