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
	DeleteLoan(ctx *gin.Context)
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
	user := ctx.MustGet("user").(*models.User)
	items, err := c.service.GetLoanSummary(user.UserID.String())
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

	err := c.service.CreateLoan(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Unexpected error`")
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *LoanController) DeleteLoan(ctx *gin.Context) {
	userId := ctx.MustGet("user").(*models.User).UserID.String()
	var input dto.DeleteLoanInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	err := c.service.DeleteLoan(userId, input.PersonName, input.IsDebt)
	if err != nil {
		if err.Error() == "loan not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, "Unexpected error")
		}
		return
	}

	ctx.Status(http.StatusOK)
}
