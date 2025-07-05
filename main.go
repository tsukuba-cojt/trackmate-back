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

	// データベースの初期化
	infra.Initialize()
	db := infra.SetupDB()

	// 支出のリポジトリ、サービス、コントローラーの初期化
	expenseRepositoty := repositories.NewExpenseRepository(db)
	expenseService := services.NewExpenseService(expenseRepositoty)
	expenseController := controllers.NewExpenseController(expenseService)

	// 支出カテゴリのリポジトリ、サービス、コントローラーの初期化
	expenseCategoryRepository := repositories.NewExpenseCategoryRepository(db)
	expenseCategoryService := services.NewExpenseCategoryService(expenseCategoryRepository)
	expenseCategoryController := controllers.NewExpenseCategoryController(expenseCategoryService)

	// 借金のリポジトリ、サービス、コントローラーの初期化
	debtRepository := repositories.NewDebtRepository(db)
	debtService := services.NewDebtService(debtRepository)
	debtController := controllers.NewDebtController(debtService)

	// 借金の人のリポジトリ、サービス、コントローラーの初期化
	debtPersonRepository := repositories.NewDebtPersonRepository(db)
	debtPersonService := services.NewDebtPersonService(debtPersonRepository)
	debtPersonController := controllers.NewDebtPersonController(debtPersonService)

	// 認証のリポジトリ、サービス、コントローラーの初期化
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	// ルーターの初期化
	r := gin.Default()

	// 支出のルーティング
	expenseRouterWithAuth := r.Group("/expenses", middlewares.AuthMiddleware(authService))
	expenseRouterWithAuth.GET("", expenseController.FindAllExpense)
	expenseRouterWithAuth.POST("", expenseController.CreateExpense)

	// 支出カテゴリのルーティング
	expenseCategoryRouterWithAuth := r.Group("/expense-categories", middlewares.AuthMiddleware(authService))
	expenseCategoryRouterWithAuth.GET("", expenseCategoryController.FindAllExpenseCategory)
	expenseCategoryRouterWithAuth.POST("", expenseCategoryController.CreateExpenseCategory)

	// 借金のルーティング
	debtRouterWithAuth := r.Group("/debts", middlewares.AuthMiddleware(authService))
	debtRouterWithAuth.GET("", debtController.FindAllDebt)
	debtRouterWithAuth.POST("", debtController.CreateDebt)

	// 借金の人のルーティング
	debtPersonRouterWithAuth := r.Group("/debt-persons", middlewares.AuthMiddleware(authService))
	debtPersonRouterWithAuth.GET("", debtPersonController.FindAllDebtPerson)
	debtPersonRouterWithAuth.POST("", debtPersonController.CreateDebtPerson)

	// 認証のルーティング
	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	// サーバーの起動
	r.Run(":8080")
}
