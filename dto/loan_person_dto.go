package dto

// 借金の相手の作成の入力データの定義
type CreateLoanPersonInput struct {
	UserID               string `json:"user_id"`
	CreateLoanPersonName string `json:"person_name" binding:"required"`
}
