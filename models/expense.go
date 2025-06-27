package models

type Expense struct {
	DbBase
	User          string `json:"user"`
	Amount        int    `json:"amount"`
	PaymentMethod string `json:"paymentmethod`
	Category      string `json:"category"`
}
