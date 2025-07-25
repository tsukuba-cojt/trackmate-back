package dto

// 借金の相手の作成の入力データの定義
type CreateLoanPersonInput struct {
	UserID               string `json:"user_id"`
	CreateLoanPersonName string `json:"person_name" binding:"required"`
}

type FindAllLoanPersonResponse struct {
	PersonID   string `json:"person_id"`
	PersonName string `json:"person_name"`
}

type DeleteLoanPersonInput struct {
	UserID   string `json:"user_id"`
	PersonID string `json:"person_id" binding:"required"`
}
