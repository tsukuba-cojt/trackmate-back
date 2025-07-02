package controllers

import (
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IExpenseCategoryController interface {
	FindAllExpenseCategory(ctx *gin.Context)
}

type ExpenseCategoryController struct {
	service services.IExpenseCategoryService
}

func NewExpenseCategoryController(service services.IExpenseCategoryService) IExpenseCategoryController {
	return &ExpenseCategoryController{service: service}
}

func (c *ExpenseCategoryController) FindAllExpenseCategory(ctx *gin.Context) {
	items, err := c.service.FindAllExpenseCategory()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
