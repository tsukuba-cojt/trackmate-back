package controllers

import (
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IDebtController interface {
	FindAllDebt(ctx *gin.Context)
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
