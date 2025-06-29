package main

import (
	"myapp/infra"
	"myapp/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	err := db.AutoMigrate(&models.Expense{}, &models.Debt{}, &models.User{}, &models.DebtPerson{}, &models.ExpenseCategory{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
