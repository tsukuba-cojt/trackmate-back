package controllers

import (
	"myapp/dto"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IExpenseController interface {
	FindAllExpense(ctx *gin.Context)
	CreateExpense(ctx *gin.Context)
}

type ExpenseController struct {
	service services.IExpenseService
}

func NewExpenseController(service services.IExpenseService) IExpenseController {
	return &ExpenseController{service: service}
}

func (c *ExpenseController) FindAllExpense(ctx *gin.Context) {
	items, err := c.service.FindAllExpense()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

func (c *ExpenseController) CreateExpense(ctx *gin.Context) {
	var input dto.CreateExpenseInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense, err := c.service.CreateExpense(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": expense})
}
