package main

import (
	"app/middlewares"
	"app/models"
	"app/services"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func Init() {
	// モデル初期化
	models.Init()

	// サービス初期化
	services.Init()

	// ミドルウェア初期化
	middlewares.Init()
}

func main() {
	// 初期化
	Init()

	// ルータ
	router := echo.New()

	// ルーティング設定
	SetupRouter(router)

	// ヘルスチェック
	router.GET("/health", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})
	
	router.Logger.Fatal(router.Start(":8080"))
}
