package main

import (
	"myapp/controllers"
	"myapp/infra"
	"myapp/repositories"
	"myapp/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	expenseRepositoty := repositories.NewExpenseRepository(db)
	expenseService := services.NewExpenseService(expenseRepositoty)
	expenseController := controllers.NewExpenseController(expenseService)

	r := gin.Default()
	expenseRouter := r.Group("/expenses")

	expenseRouter.GET("", expenseController.FindAllExpense)
	expenseRouter.POST("", expenseController.CreateExpense)

	expenseCategoryRepository := repositories.NewExpenseCategoryRepository(db)
	expenseCategoryService := services.NewExpenseCategoryService(expenseCategoryRepository)
	expenseCategoryController := controllers.NewExpenseCategoryController(expenseCategoryService)
	expenseCategoryRouter := r.Group("/expense-categories")
	expenseCategoryRouter.GET("", expenseCategoryController.FindAllExpenseCategory)

	debtRepository := repositories.NewDebtRepository(db)
	debtService := services.NewDebtService(debtRepository)
	debtController := controllers.NewDebtController(debtService)
	debtRouter := r.Group("/debts")
	debtRouter.GET("", debtController.FindAllDebt)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)
	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	r.Run(":8080")
}
