package main

import (
	"app/controllers"
	"app/middlewares"
	"app/utils"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter(router *echo.Echo) {
	// logger 設定
	router.Use(middleware.Logger())

	router.Use(session.Middleware(utils.SessionStore))

	// link開始
	router.POST("/startlink", controllers.StartLink,middlewares.RequireAuth)

	// 実際にリンク
	router.GET("/aclink", controllers.AcLink)

	// コールバック
	router.GET("/link/callback", controllers.Callback)

	// linkステータス
	router.GET("/link/status", controllers.LinkStatus,middlewares.RequireAuth)

	// 同意エンドポイント
	router.POST("/terms/accept", controllers.AcceptTerms,middlewares.RequireAuth)

	// hello world
	router.GET("/hello", controllers.Hello,middlewares.RequireAuth)
}