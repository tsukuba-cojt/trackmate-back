package controllers

import (
	"myapp/dto"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// インターフェースの定義
type IAuthController interface {
	Signup(ctx *gin.Context)
	Login(ctx *gin.Context)
}

// コントローラーの定義
type AuthController struct {
	service services.IAuthService
}

// コンストラクタの定義
func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

// ユーザーを作成する関数の定義
func (c *AuthController) Signup(ctx *gin.Context) {
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := c.service.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.Status(http.StatusCreated)
}

// ユーザーをログインさせる関数の定義
func (c *AuthController) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
