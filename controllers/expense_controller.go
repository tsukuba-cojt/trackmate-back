package controllers

import (
	"myapp/dto"      // dto.CreateExpenseInput で使用するため必要
	"myapp/models"   // models.User で使用するため必要
	"myapp/services" // services.IExpenseService で使用するため必要
	"net/http"

	"github.com/gin-gonic/gin"
)

// インターフェースの定義
type IExpenseController interface {
	GetExpenseSummary(ctx *gin.Context)
	CreateExpense(ctx *gin.Context)
	DeleteExpense(ctx *gin.Context)
}

// コントローラーの定義
type ExpenseController struct {
	expenseService services.IExpenseService
	summaryFacade  services.ISummaryFacade
}

// コンストラクタの定義
func NewExpenseController(
	expenseService services.IExpenseService,
	summaryFacade services.ISummaryFacade,
) IExpenseController {
	return &ExpenseController{
		expenseService: expenseService,
		summaryFacade:  summaryFacade,
	}
}

// GetExpenseSummary はユーザーごとの支出のサマリーを取得する関数の定義
func (c *ExpenseController) GetExpenseSummary(ctx *gin.Context) {
	// 認証ミドルウェアからユーザー情報を取得
	user := ctx.MustGet("user").(*models.User)
	userId := user.UserID.String()

	date := ctx.Query("date")
	if date != "" {
		// クエリパラメータに日付があった時はその日の支出　の一覧を返す
		summary, err := c.summaryFacade.GetExpenseSummaryByDate(userId, date)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": summary})
		return
	} else {
		// クエリパラメータが空の場合は当月のサマリーを返す
		summary, err := c.summaryFacade.GetExpenseSummary(userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": summary})
		return
	}
}

// CreateExpense は支出を作成する関数の定義
func (c *ExpenseController) CreateExpense(ctx *gin.Context) {
	var input dto.CreateExpenseInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user := ctx.MustGet("user").(*models.User)
	input.UserID = user.UserID.String() // DTOにUserIDを設定

	// 個別サービスを使用（ファサードは統合処理のみ）
	_, err := c.expenseService.CreateExpense(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense"})
		return
	}

	ctx.Status(http.StatusCreated)
}

// DeleteExpense は支出を削除する関数の定義
func (c *ExpenseController) DeleteExpense(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	userId := user.UserID.String()

	var input dto.DeleteExpenseInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	expenseId := input.ExpenseID

	// 個別サービスを使用（ファサードは統合処理のみ）
	err := c.expenseService.DeleteExpense(userId, expenseId)
	if err != nil {
		if err.Error() == "expense not found or not authorized to delete" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.Status(http.StatusOK)
}
