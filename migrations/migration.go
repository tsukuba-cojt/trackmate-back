package main

import (
	"myapp/infra"
	"myapp/models"
)

func main() {
	// データベースの初期化
	infra.Initialize()
	db := infra.SetupDB()

	// マイグレーションの実行
	err := db.AutoMigrate(&models.Expense{}, &models.Loan{}, &models.User{}, &models.LoanPerson{}, &models.ExpenseCategory{}, &models.Budget{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
