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

	// 支出のリポジトリ、サービス、コントローラーの初期化
	expenseRepository := repositories.NewExpenseRepository(db)
	expenseService := services.NewExpenseService(expenseRepository)
	summaryFacade := services.NewSummaryFacade(expenseRepository, budgetRepository, loanRepository)
	expenseController := controllers.NewExpenseController(expenseService, summaryFacade)

	// 認証のリポジトリ、サービス、コントローラーの初期化
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	// ルーターの初期化
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://mast23mc.net"}, // フロントエンドのURLを明示的に許可
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // プリフライトリクエストをキャッシュ
	}))

	api := r.Group("/api")

	// 支出のルーティング
	expenseRouterWithAuth := api.Group("/expenses", middlewares.AuthMiddleware(authService))
	expenseRouterWithAuth.GET("", expenseController.GetExpenseSummary)
	expenseRouterWithAuth.POST("", expenseController.CreateExpense)
	expenseRouterWithAuth.DELETE("", expenseController.DeleteExpense)

	// 支出カテゴリのルーティング
	expenseCategoryRouterWithAuth := api.Group("/categories", middlewares.AuthMiddleware(authService))
	expenseCategoryRouterWithAuth.GET("", expenseCategoryController.GetExpenseCategorySummary)
	expenseCategoryRouterWithAuth.POST("", expenseCategoryController.CreateExpenseCategory)
	expenseCategoryRouterWithAuth.DELETE("", expenseCategoryController.DeleteExpenseCategory)

	// 借金のルーティング
	loanRouterWithAuth := api.Group("/loan", middlewares.AuthMiddleware(authService))
	loanRouterWithAuth.GET("", loanController.GetLoanSummary)
	loanRouterWithAuth.POST("", loanController.CreateLoan)
	loanRouterWithAuth.DELETE("", loanController.DeleteLoan)

	// 借金の人のルーティング
	loanPersonRouterWithAuth := api.Group("/person", middlewares.AuthMiddleware(authService))
	loanPersonRouterWithAuth.GET("", loanPersonController.FindAllLoanPerson)
	loanPersonRouterWithAuth.POST("", loanPersonController.CreateLoanPerson)
	loanPersonRouterWithAuth.DELETE("", loanPersonController.DeleteLoanPerson)

	// 予算のルーティング
	budgetRouterWithAuth := api.Group("/budget", middlewares.AuthMiddleware(authService))
	budgetRouterWithAuth.POST("", budgetController.CreateBudget)

	// 認証のルーティング
	authRouter := api.Group("/auth")
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	// サーバーの起動
	r.Run(":3401")
}
