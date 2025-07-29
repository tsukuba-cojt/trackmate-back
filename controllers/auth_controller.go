package controllers

import (
	"myapp/dto"
	"myapp/services"
	"net/http"
	"time"

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
		return
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

	// サービス層のLoginメソッドを呼び出し、JWT文字列へのポインタを取得
	tokenPtr, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	// tokenPtr は *string 型なので、Cookieに設定するために逆参照して string 型にする
	var stringToken string
	if tokenPtr != nil {
		stringToken = *tokenPtr
	} else {
		// tokenPtr が nil の場合はエラーとするか、適切に処理
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Login successful but token is nil"})
		return
	}

	// Cookie設定
	cookie := new(http.Cookie)
	cookie.Name = "trackmate_auth_token"
	cookie.Value = stringToken
	cookie.Expires = time.Now().Add(6 * time.Hour)
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Path = "/"
	cookie.Secure = true

	// Cookieをレスポンスヘッダーに設定
	ctx.SetCookie(cookie.Name, cookie.Value, int(cookie.Expires.Unix()), cookie.Path, "", cookie.Secure, cookie.HttpOnly)

	ctx.JSON(http.StatusOK, gin.H{"token": tokenPtr}) // クライアントにはポインタのままJWTを返す
}
