package controllers

import (
	"myapp/dto"
	"myapp/models"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// インターフェースの定義
type IBudgetController interface {
	CreateBudget(ctx *gin.Context)
}

// コントローラーの定義
type BudgetController struct {
	service services.IBudgetService
}

// コンストラクタの定義
func NewBudgetController(service services.IBudgetService) IBudgetController {
	return &BudgetController{service: service}
}

// 予算を作成する関数の定義
func (c *BudgetController) CreateBudget(ctx *gin.Context) {
	var input dto.CreateBudgetInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	user := ctx.MustGet("user").(*models.User)
	input.UserID = user.UserID.String()

	err := c.service.CreateBudget(input)
	if err != nil {
		if err.Error() == "budget already exists" {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Budget already exists"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}

	ctx.Status(http.StatusCreated)
}
