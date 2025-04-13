package main

import (
	"app/controllers"
	"app/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter(router *echo.Echo) {
	// logger 設定
	router.Use(middleware.Logger())

	// hello world
	router.GET("/hello", controllers.Hello,middlewares.RequireAuth)
}