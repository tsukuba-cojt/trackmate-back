package dto

type CreateDebtInput struct {
	UserID       string `json:"user_id"`
	DebtPersonID string `json:"debt_person_id" binding:"required"`
	IsBorrow     bool   `json:"is_borrow" binding:"required"`
	DebtDate     string `json:"debt_date" binding:"required"`
	DebtAmount   int    `json:"amount" binding:"required"`
	Description  string `json:"description" binding:"required"`
}
