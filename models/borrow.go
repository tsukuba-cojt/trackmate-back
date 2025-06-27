package models

type Borrow struct {
	DbBase
	User   string `json:"user"`
	People string `json:"people"`
	Amount int    `json:"amount"`
}
