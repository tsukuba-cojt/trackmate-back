package main

import (
	"myapp/controllers"
	"myapp/infra"
	"myapp/repositories"
	"myapp/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	expenseRepositoty := repositories.NewExpenseRepository(db)
	expenseService := services.NewExpenseService(expenseRepositoty)
	expenseController := controllers.NewExpenseController(expenseService)

	r := gin.Default()
	expenseRouter := r.Group("/expenses")

	expenseRouter.GET("", expenseController.FindAll)

	r.Run(":8080")
}
