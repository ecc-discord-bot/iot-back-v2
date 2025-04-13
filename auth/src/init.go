package main

import (
	"auth/controllers"
	"auth/middlewares"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (temp *TemplateRenderer) Render(writer io.Writer, name string, data interface{}, ctx echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = ctx.Echo().Reverse
	}

	return temp.templates.ExecuteTemplate(writer, name, data)
}

func SetupRouter(router *echo.Echo) {
	// logger 設定
	router.Use(middleware.Logger())

	// テンプレート
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	// レンダラー
	router.Renderer = renderer

	// ルーティング設定
	// ベーシックユーザーグループ
	basicg := router.Group("/basic")
	{
		basicg.POST("/signup", controllers.CreateBasicUser)
		basicg.POST("/login", controllers.LoginBasicUser)
	}

	// 情報を取得する
	router.GET("/me", controllers.GetMe, middlewares.RequireAuth)

	// token を取得する
	router.GET("/token", controllers.GetToken, middlewares.RequireAuth)

	// ログアウト
	router.POST("/logout",controllers.Logout,middlewares.RequireAuth)

	// oauth グループ
	oauthg := router.Group("/oauth")
	{
		oauthg.GET("/:provider",controllers.StartOauth)
		oauthg.GET("/:provider/callback",controllers.CallbackOauth)
	}
}