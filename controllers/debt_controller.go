package controllers

import (
	"myapp/dto"
	"myapp/models"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// インターフェースの定義
type IDebtController interface {
	FindAllDebt(ctx *gin.Context)
	CreateDebt(ctx *gin.Context)
}

// コントローラーの定義
type DebtController struct {
	service services.IDebtService
}

// コンストラクタの定義
func NewDebtController(service services.IDebtService) IDebtController {
	return &DebtController{service: service}
}

// ユーザーごとの全ての借金を取得する関数の定義
func (c *DebtController) FindAllDebt(ctx *gin.Context) {
	items, err := c.service.FindAllDebt()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// 借金を作成する関数の定義
func (c *DebtController) CreateDebt(ctx *gin.Context) {
	var input dto.CreateDebtInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := ctx.MustGet("user").(*models.User)
	input.UserID = user.UserID.String()

	debt, err := c.service.CreateDebt(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": debt})
}
