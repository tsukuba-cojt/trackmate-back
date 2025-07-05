package dto

// 借金の相手の作成の入力データの定義
type CreateDebtPersonInput struct {
	UserID               string `json:"user_id"`
	CreateDebtPersonName string `json:"debt_person_name" binding:"required"`
}
