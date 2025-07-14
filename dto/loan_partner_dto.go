package dto

// 借金の相手の作成の入力データの定義
type CreateLoanPartnerInput struct {
	UserID                string `json:"user_id"`
	CreateLoanPartnerName string `json:"loan_partner_name" binding:"required"`
}
