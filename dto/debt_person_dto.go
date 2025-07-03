package dto

type CreateDebtPersonInput struct {
	UserID               string `json:"user_id"`
	CreateDebtPersonName string `json:"debt_person_name" binding:"required"`
}
