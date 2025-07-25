package dto

type CreateBudgetInput struct {
	Budget int    `json:"budget"`
	Date   string `json:"date"`
	UserID string `json:"user_id"`
}
