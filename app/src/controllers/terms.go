package controllers

import (
	"app/logger"
	"app/middlewares"
	"app/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccepterTerms struct {
	UserName string `json:"UserName"`
	Class    string `json:"UserClass"`
}

func AcceptTerms(ctx echo.Context) error {
	// claimを取得
	claim := ctx.Get("claim").(middlewares.AccessTokenClaim)

	// bindする
	var terms AccepterTerms
	if err := ctx.Bind(&terms); err != nil {
		logger.PrintErr(err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	logger.Println(terms)

	// 認証を完了
	if err := services.AcceptTerms(claim, services.AccepterTerms{
		UserName: terms.UserName,
		Class:    terms.Class,
	}); err != nil {
		logger.PrintErr(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to accept terms"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "terms accepted"})
}
