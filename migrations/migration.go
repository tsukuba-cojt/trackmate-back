package main

import (
	"fmt"
	"myapp/infra"
	"myapp/models"
	"time"

	"github.com/google/uuid"
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

	// --- ここから初期データ挿入 ---
	userID := uuid.New()
	fmt.Println(userID)
	user := models.User{
		UserID:    userID,
		Email:     "test@example.com",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&user)

	expenseCategoryID := uuid.New()
	fmt.Println(expenseCategoryID)
	expenseCategory := models.ExpenseCategory{
		ExpenseCategoryID:   expenseCategoryID,
		UserID:              userID,
		ExpenseCategoryName: "食費",
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
	db.Create(&expenseCategory)
	// --- ここまで ---
}
