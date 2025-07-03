package main

import (
	"myapp/controllers"
	"myapp/infra"
	"myapp/middlewares"
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

	expenseCategoryRepository := repositories.NewExpenseCategoryRepository(db)
	expenseCategoryService := services.NewExpenseCategoryService(expenseCategoryRepository)
	expenseCategoryController := controllers.NewExpenseCategoryController(expenseCategoryService)

	debtRepository := repositories.NewDebtRepository(db)
	debtService := services.NewDebtService(debtRepository)
	debtController := controllers.NewDebtController(debtService)

	debtPersonRepository := repositories.NewDebtPersonRepository(db)
	debtPersonService := services.NewDebtPersonService(debtPersonRepository)
	debtPersonController := controllers.NewDebtPersonController(debtPersonService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	expenseRouterWithAuth := r.Group("/expenses", middlewares.AuthMiddleware(authService))

	expenseRouterWithAuth.GET("", expenseController.FindAllExpense)
	expenseRouterWithAuth.POST("", expenseController.CreateExpense)

	expenseCategoryRouterWithAuth := r.Group("/expense-categories", middlewares.AuthMiddleware(authService))
	expenseCategoryRouterWithAuth.GET("", expenseCategoryController.FindAllExpenseCategory)
	expenseCategoryRouterWithAuth.POST("", expenseCategoryController.CreateExpenseCategory)

	debtRouterWithAuth := r.Group("/debts", middlewares.AuthMiddleware(authService))
	debtRouterWithAuth.GET("", debtController.FindAllDebt)

	debtPersonRouterWithAuth := r.Group("/debt-persons", middlewares.AuthMiddleware(authService))
	debtPersonRouterWithAuth.GET("", debtPersonController.FindAllDebtPerson)
	debtPersonRouterWithAuth.POST("", debtPersonController.CreateDebtPerson)

	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	r.Run(":8080")
}
