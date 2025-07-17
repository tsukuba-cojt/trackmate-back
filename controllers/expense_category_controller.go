package controllers

import (
	"myapp/dto"
	"myapp/models"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// インターフェースの定義
type IExpenseCategoryController interface {
	GetExpenseCategory(ctx *gin.Context)
	CreateExpenseCategory(ctx *gin.Context)
}

// コントローラーの定義
type ExpenseCategoryController struct {
	service services.IExpenseCategoryService
}

// コンストラクタの定義
func NewExpenseCategoryController(service services.IExpenseCategoryService) IExpenseCategoryController {
	return &ExpenseCategoryController{service: service}
}

// ユーザーごとの全ての支出カテゴリを取得する関数の定義
func (c *ExpenseCategoryController) GETExpenseCategory(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	userId := user.UserID.String()
	items, err := c.service.FindAllExpenseCategory(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// 支出カテゴリを作成する関数の定義
func (c *ExpenseCategoryController) CreateExpenseCategory(ctx *gin.Context) {
	var input dto.CreateExpenseCategoryInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := ctx.MustGet("user").(*models.User)
	input.UserID = user.UserID.String()

	expenseCategory, err := c.service.CreateExpenseCategory(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": expenseCategory})
}
