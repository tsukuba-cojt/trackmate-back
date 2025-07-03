package controllers

import (
	"myapp/dto"
	"myapp/models"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IDebtPersonController interface {
	FindAllDebtPerson(ctx *gin.Context)
	CreateDebtPerson(ctx *gin.Context)
}

type DebtPersonController struct {
	service services.IDebtPersonService
}

func NewDebtPersonController(service services.IDebtPersonService) IDebtPersonController {
	return &DebtPersonController{service: service}
}

func (c *DebtPersonController) FindAllDebtPerson(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	userId := user.UserID.String()
	items, err := c.service.FindAllDebtPerson(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

func (c *DebtPersonController) CreateDebtPerson(ctx *gin.Context) {
	var input dto.CreateDebtPersonInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := ctx.MustGet("user").(*models.User)
	input.UserID = user.UserID.String()

	debtPerson, err := c.service.CreateDebtPerson(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": debtPerson})
}
