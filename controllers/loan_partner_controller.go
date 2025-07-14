package controllers

import (
	"myapp/dto"
	"myapp/models"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// インターフェースの定義
type ILoanPartnerController interface {
	FindAllLoanPartner(ctx *gin.Context)
	CreateLoanPartner(ctx *gin.Context)
}

// コントローラーの定義
type LoanPartnerController struct {
	service services.ILoanPartnerService
}

// コンストラクタの定義
func NewLoanPartnerController(service services.ILoanPartnerService) ILoanPartnerController {
	return &LoanPartnerController{service: service}
}

// ユーザーごとの全ての借金の相手を取得する関数の定義
func (c *LoanPartnerController) FindAllLoanPartner(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	userId := user.UserID.String()
	items, err := c.service.FindAllLoanPartner(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// 借金の相手を作成する関数の定義
func (c *LoanPartnerController) CreateLoanPartner(ctx *gin.Context) {
	var input dto.CreateLoanPersonInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := ctx.MustGet("user").(*models.User)
	input.UserID = user.UserID.String()

	loanPartner, err := c.service.CreateLoanPartner(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": loanPartner})
}
