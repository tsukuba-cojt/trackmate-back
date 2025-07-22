package controllers

import (
	"myapp/dto"
	"myapp/models"
	"myapp/services"
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

// ユーザーごとの全ての支出を取得する関数の定義
func (c *ExpenseController) GetExpenseSummary(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	userId := user.UserID.String()
	items, err := c.service.FindAllExpense(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// 支出を作成する関数の定義
func (c *ExpenseController) CreateExpense(ctx *gin.Context) {
	var input dto.CreateExpenseInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	user := ctx.MustGet("user").(*models.User)
	input.UserID = user.UserID.String()

	_, err := c.service.CreateExpense(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.Status(http.StatusOK)
}


func (c *ExpenseController) DeleteExpense(ctx *gin.Context) {
	expenseId := ctx.Param("id")
	if expenseId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Expense ID is required"})
		return
	}

	err := c.service.DeleteExpense(expenseId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.Status(http.StatusNoContent)
}