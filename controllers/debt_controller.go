package controllers

import (
	"myapp/dto"
	"myapp/models"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IDebtController interface {
	FindAllDebt(ctx *gin.Context)
	CreateDebt(ctx *gin.Context)
}

type DebtController struct {
	service services.IDebtService
}

func NewDebtController(service services.IDebtService) IDebtController {
	return &DebtController{service: service}
}

func (c *DebtController) FindAllDebt(ctx *gin.Context) {
	items, err := c.service.FindAllDebt()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

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
