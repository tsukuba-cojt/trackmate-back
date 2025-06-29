package controllers

import (
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IExpenseController interface {
	FindAll(ctx *gin.Context)
}

type ExpenseController struct {
	service services.IExpenseService
}

func NewExpenseController(service services.IExpenseService) IExpenseController {
	return &ExpenseController{service: service}
}

func (c *ExpenseController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
