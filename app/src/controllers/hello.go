package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Hello, World!"})
}