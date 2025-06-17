package main

import (
	//"net/http"
	"os"
	"log"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	//"database/sql"
	_"github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Expense struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey"`
	User string `json:"user"`
	Amount int `json:"amount"`
	PaymentMethod string `json:"paymentmethod`
	Category string `json:"category"`
}

type 

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB = db

	err = DB.AutoMigrate(&Expense{})
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}
	log.Println("Database connected successfully!")

}


func createExpenses(c *gin.Context) {
	var newExpense Expense
	newExpense.ID = uuid.New()
	if err := c.ShouldBindJSON(&newExpense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Create(&newExpense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newExpense)
}


func getExpenses(c *gin.Context) {
	var expenses []Expense
	DB.Find(&expenses)
	c.JSON(http.StatusOK, expenses)
}

func main() {
	r := gin.Default()

	ConnectDB()

	r.POST("/expenses", createExpenses)

	r.GET("/expenses", getExpenses)

	r.Run(":8080")
}