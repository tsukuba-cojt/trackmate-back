package main

import (
	"myapp/controllers"
	"myapp/infra"
	"myapp/middlewares"
	"myapp/repositories"
	"myapp/services"
	"time"

	"github.com/gin-contrib/cors"
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
	loanRepository := repositories.NewLoanRepository(db)
	loanService := services.NewLoanService(loanRepository)
	loanController := controllers.NewLoanController(loanService)

	// 借金の人のリポジトリ、サービス、コントローラーの初期化
	loanPersonRepository := repositories.NewLoanPersonRepository(db)
	loanPersonService := services.NewLoanPersonService(loanPersonRepository)
	loanPersonController := controllers.NewLoanPersonController(loanPersonService)

	// 予算のリポジトリ、サービス、コントローラーの初期化
	budgetRepository := repositories.NewBudgetRepository(db)
	budgetService := services.NewBudgetService(budgetRepository)
	budgetController := controllers.NewBudgetController(budgetService)

	// 認証のリポジトリ、サービス、コントローラーの初期化
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	// ルーターの初期化
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3400"}, // フロントエンドのURLを明示的に許可
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // プリフライトリクエストをキャッシュ
	}))

	// 支出のルーティング
	expenseRouterWithAuth := r.Group("/expenses", middlewares.AuthMiddleware(authService))
	expenseRouterWithAuth.GET("", expenseController.FindAllExpense)
	expenseRouterWithAuth.POST("", expenseController.CreateExpense)

	// 支出カテゴリのルーティング
	expenseCategoryRouterWithAuth := r.Group("/categories", middlewares.AuthMiddleware(authService))
	expenseCategoryRouterWithAuth.GET("", expenseCategoryController.GetExpenseCategorySummary)
	expenseCategoryRouterWithAuth.POST("", expenseCategoryController.CreateExpenseCategory)
	expenseCategoryRouterWithAuth.DELETE("", expenseCategoryController.DeleteExpenseCategory)

	// 借金のルーティング
	loanRouterWithAuth := r.Group("/loan", middlewares.AuthMiddleware(authService))
	loanRouterWithAuth.GET("", loanController.GetLoanSummary)
	loanRouterWithAuth.POST("", loanController.CreateLoan)
	loanRouterWithAuth.DELETE("", loanController.DeleteLoan)

	// 借金の人のルーティング
	loanPersonRouterWithAuth := r.Group("/person", middlewares.AuthMiddleware(authService))
	loanPersonRouterWithAuth.GET("", loanPersonController.FindAllLoanPerson)
	loanPersonRouterWithAuth.POST("", loanPersonController.CreateLoanPerson)
	loanPersonRouterWithAuth.DELETE("", loanPersonController.DeleteLoanPerson)

	// 予算のルーティング
	budgetRouterWithAuth := r.Group("/budget", middlewares.AuthMiddleware(authService))
	budgetRouterWithAuth.POST("", budgetController.CreateBudget)

	// 認証のルーティング
	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	// サーバーの起動
	r.Run(":8000")
}
