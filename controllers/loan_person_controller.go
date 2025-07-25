package controllers

import (
	"myapp/dto"
	"myapp/models"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// インターフェースの定義
type ILoanPersonController interface {
	FindAllLoanPerson(ctx *gin.Context)
	CreateLoanPerson(ctx *gin.Context)
	DeleteLoanPerson(ctx *gin.Context)
}

// コントローラーの定義
type LoanPersonController struct {
	service services.ILoanPersonService
}

// コンストラクタの定義
func NewLoanPersonController(service services.ILoanPersonService) ILoanPersonController {
	return &LoanPersonController{service: service}
}

// ユーザーごとの全ての借金の相手を取得する関数の定義
func (c *LoanPersonController) FindAllLoanPerson(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	userId := user.UserID.String()
	items, err := c.service.FindAllLoanPerson(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// 借金の相手を作成する関数の定義
func (c *LoanPersonController) CreateLoanPerson(ctx *gin.Context) {
	var input dto.CreateLoanPersonInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := ctx.MustGet("user").(*models.User)
	input.UserID = user.UserID.String()

	err := c.service.CreateLoanPerson(input)
	if err != nil {
		if err.Error() == "loan person already exists" {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Loan person already exists"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
			return
		}
	}

	ctx.Status(http.StatusCreated)
}

func (c *LoanPersonController) DeleteLoanPerson(ctx *gin.Context) {
	var input dto.DeleteLoanPersonInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	user := ctx.MustGet("user").(*models.User)
	input.UserID = user.UserID.String()

	err := c.service.DeleteLoanPerson(input)
	if err != nil {
		if err.Error() == "loan person not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Loan person not found"})
			return
		} else if err.Error() == "loan person is referenced by loans" {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Loan person is referenced by loans"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
			return
		}
	}

	ctx.Status(http.StatusOK)
}
