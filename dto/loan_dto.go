package dto

// 借金の作成の入力データの定義
type CreateLoanInput struct {
	UserID       string `json:"user_id"`
	LoanPersonID string `json:"person_id" binding:"required"`
	IsDebt       bool   `json:"is_debt" binding:"required"`
	LoanDate     string `json:"date" binding:"required"`
	LoanAmount   int    `json:"amount" binding:"required"`
	Description  string `json:"description"`
}

type LoanSummaryResponse struct {
	PersonName string                `json:"person_name"`
	IsDebt     bool                  `json:"is_debt"`
	SumAmount  int                   `json:"sum_amount"`
	History    []LoanHistoryResponse `json:"history"`
}

type LoanHistoryResponse struct {
	Date   string `json:"date"`
	Amount int    `json:"amount"`
}
