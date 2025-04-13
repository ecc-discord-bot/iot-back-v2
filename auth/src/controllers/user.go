package controllers

import (
	"auth/models"
	"auth/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMe(ctx echo.Context) error {
	// セッションを取得
	session, ok := ctx.Get("session").(*models.Session)

	// エラー処理
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	// 自身の情報を取得
	user, err := services.GetMe(session.UserID)

	// エラー処理
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, user)
}