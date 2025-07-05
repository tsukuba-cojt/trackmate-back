package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

// インターフェースの定義
type IAuthRepository interface {
	CreateUser(user models.User) error
	FindUser(email string) (*models.User, error)
}

// リポジトリの定義
type AuthRepository struct {
	db *gorm.DB
}

// コンストラクタの定義
func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db: db}
}

// ユーザーを作成する関数の定義
func (r *AuthRepository) CreateUser(user models.User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ユーザーを取得する関数の定義
func (r *AuthRepository) FindUser(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
