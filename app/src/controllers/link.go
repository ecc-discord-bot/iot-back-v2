package controllers

import (
	"app/logger"
	"app/middlewares"
	"app/oauth2"
	"app/services"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func StartLink(ctx echo.Context) error {
	// tokenを取得
	token := ctx.Get("token")

	// セッションを取得
	sess, err := session.Get("linkSession", ctx)
	if err != nil {
		logger.PrintErr(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to get session"})
	}

	// セッションを設定
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   15 * 60,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	// サンプルvalueを設定
	sess.Values["token"] = token
	
	// セッションを保存
	err = sess.Save(ctx.Request(), ctx.Response())

	if err != nil {
		logger.PrintErr(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to save session"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "start link"})
}

func AcLink(ctx echo.Context) error {
	// セッションを取得
	sess, err := session.Get("linkSession", ctx)
	if err != nil {
		logger.PrintErr(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to get session"})
	}

	// サンプルvalueを取得
	sampleValue := sess.Values["sampleValue"]

	logger.Println(sampleValue)

	// 認証を開始
	return oauth2.StartOauth(ctx, oauth2.OauthArgs{ProviderName: "microsoftonline"})
}

func Callback(ctx echo.Context) error {
	// セッションを取得
	sess, err := session.Get("linkSession", ctx)
	if err != nil {
		logger.PrintErr(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to get session"})
	}	

	// トークンを取得
	token := sess.Values["token"]

	// 認証を完了
	response, err := oauth2.CallbackOauth(ctx, "microsoftonline")

	// エラー処理
	if err != nil {
		logger.PrintErr(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to get session"})
	}

	// バリデーションを実行する
	claim, err := middlewares.ValidateToken(token.(string))

	// エラー処理
	if err != nil {
		logger.PrintErr(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to get session"})
	}

	// プロバイダを検証する
	if claim.ProvCode != "discord" {
		logger.PrintErr("ProvCode is not discord")
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Discord 以外ではリンクできません"})
	}

	// メールアドレスを検証する (@ecc.ac.jp で終わっているか)
	if response.User.Email == "" || !strings.HasSuffix(response.User.Email, "@ecc.ac.jp") {
		logger.PrintErr("Email is not ecc.ac.jp")
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "ecc.ac.jp で終わるメールアドレスしかリンクできません"})
	}

	// リンクを実行する
	err = services.LinkAccount(claim, response.User)

	// エラー処理
	if err != nil {
		logger.PrintErr(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to link account"})
	}

	// リダイレクト
	return ctx.Redirect(http.StatusSeeOther, "/statics/")
}

func LinkStatus(ctx echo.Context) error {
	// claimを取得
	claim := ctx.Get("claim").(middlewares.AccessTokenClaim)

	// ユーザーを取得する
	user, err := services.GetDiscordUser(claim.ProvUid)

	// エラー処理
	if err != nil {
		logger.PrintErr(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to get user"})
	}

	return ctx.JSON(http.StatusOK, user)
}