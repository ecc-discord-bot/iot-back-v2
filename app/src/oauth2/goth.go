package oauth2

import (
	"app/utils"
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/microsoftonline"
)

func InitGothic() {
	gothic.Store = utils.SessionStore

	UseProviders()
}

type OauthArgs struct {
	ProviderName string // プロバイダー名
}

// 認証を開始するメソッド
func StartOauth(ctx echo.Context, args OauthArgs) error {
	// リクエストを変更
	ctx.SetRequest(contextWithProviderName(ctx, args.ProviderName))

	// リクエスト取得
	request := ctx.Request()
	response := ctx.Response()

	// 認証開始
	gothic.BeginAuthHandler(response.Writer, request)

	return nil
}

// 認証を完了
type OauthResponse struct {
	User goth.User
}

func CallbackOauth(ctx echo.Context, providerName string) (OauthResponse, error) {
	request := contextWithProviderName(ctx, providerName)

	// リクエスト変更
	ctx.SetRequest(request)

	// 認証を完了
	user, err := gothic.CompleteUserAuth(ctx.Response().Writer, request)

	// エラー処理
	if err != nil {
		return OauthResponse{}, err
	}

	return OauthResponse{User: user}, nil
}

// コンテキストを設定
func contextWithProviderName(ctx echo.Context, providerName string) *http.Request {
	return ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), "provider", providerName))
}

func UseProviders() {
	goth.UseProviders(microsoftonline.New(os.Getenv("MicrosoftClientId"), os.Getenv("MicrosoftClientSecret"), os.Getenv("MicrosoftCallbackURL")))
}
