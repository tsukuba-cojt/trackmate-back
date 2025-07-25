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
	service services.IExpenseService
}

// コンストラクタの定義
func NewExpenseController(service services.IExpenseService) IExpenseController {
	return &ExpenseController{service: service}
}

// GetExpenseSummary はユーザーごとの支出のサマリーを取得する関数の定義
// サービス層のExpenseSummaryを呼び出すように変更しました。
func (c *ExpenseController) GetExpenseSummary(ctx *gin.Context) {
	// 認証ミドルウェアからユーザー情報を取得
	user := ctx.MustGet("user").(*models.User)
	userId := user.UserID.String()

	summary, err := c.service.ExpenseSummary(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error retrieving expense summary"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": summary})
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

	_, err := c.service.CreateExpense(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense"})
		return
	}

	ctx.Status(http.StatusOK)
}

// DeleteExpense は支出を削除する関数の定義
func (c *ExpenseController) DeleteExpense(ctx *gin.Context) {

	user := ctx.MustGet("user").(*models.User)
	userId := user.UserID.String()

	expenseId := ctx.Param("id")
	if expenseId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Expense ID is required"})
		return
	}

	err := c.service.DeleteExpense(userId, expenseId)
	if err != nil {
		if err.Error() == "expense not found or not authorized to delete" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expense"})
		return
	}

	ctx.Status(http.StatusNoContent)
}