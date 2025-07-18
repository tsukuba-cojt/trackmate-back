package controllers

import (
	"myapp/dto"
	"myapp/models"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// インターフェースの定義
type ILoanController interface {
	GetLoanSummary(ctx *gin.Context)
	CreateLoan(ctx *gin.Context)
}

// コントローラーの定義
type LoanController struct {
	service services.ILoanService
}

// コンストラクタの定義
func NewLoanController(service services.ILoanService) ILoanController {
	return &LoanController{service: service}
}

// ユーザーごとの全ての借金を取得する関数の定義
func (c *LoanController) GetLoanSummary(ctx *gin.Context) {
	items, err := c.service.GetLoanSummary()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// 借金を作成する関数の定義
func (c *LoanController) CreateLoan(ctx *gin.Context) {
	var input dto.CreateLoanInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	user := ctx.MustGet("user").(*models.User)
	input.UserID = user.UserID.String()

	loan, err := c.service.CreateLoan(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
