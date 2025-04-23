package controllers

import (
	"app/logger"
	"app/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SyncData(ctx echo.Context) error {
	// サービスを呼び出す
	if err := services.SyncData(); err != nil {
		logger.PrintErr(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to sync data"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "sync data"})
}