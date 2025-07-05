package infra

import (
	"log"

	"github.com/joho/godotenv"
)

// 環境変数の読み込みを行う関数の定義
func Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
