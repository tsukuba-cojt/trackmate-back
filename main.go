package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	ID		uint	`gorm:"primaryKey" json:"id"`
	Title	string 	`json:"title"`
	Completed	bool `json:"completed"`
}

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})
}

func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func getTodos(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func main() {
	r := gin.Default()

	initDB()

	r.POST("/todos", createTodo)

	r.GET("/todos", getTodos)

	r.Run(":8080")
}