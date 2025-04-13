package middlewares

import (
	"auth/logger"
	"auth/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 認証ミドルウェア
func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// ヘッダからトークンを取得
		token := ctx.Request().Header.Get("Authorization")
		if token == "" {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
		}

		// トークンを検証
		session, err := services.GetSession(token)
		if err != nil {
			logger.PrintErr(err)
			return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
		}

		// セッションを設定
		ctx.Set("session", session)

		// 認証処理
		return next(ctx)
	}
}