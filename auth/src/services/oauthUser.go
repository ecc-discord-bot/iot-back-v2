package services

import (
	"auth/models"
	"auth/utils"
)

type OauthUserArgs struct {
	Name           string // ユーザー名
	Email          string // メールアドレス
	ProviderCode   string // 認証プロバイダコード
	ProviderUserID string // 認証プロバイダユーザーID
	RemoteIP       string // IPアドレス
	UserAgent      string // User-Agent
}

// Oauthユーザーを作成する
func LoginOauthUser(args OauthUserArgs) (string, error) {
	// UUID を生成
	uid := utils.GenID()

	// 現在時刻を取得
	now := utils.NowTime()

	// ユーザーを取得する
	user, err := models.GetUserByEmail(args.Email)

	// エラー処理
	if err == nil {
		// ユーザーが取得できた時
		// セッションを追加する
		token, err := NewSession(SessionArgs{
			UserID:    user.UserID,
			RemoteIP:  args.RemoteIP,
			UserAgent: args.UserAgent,
		})

		return token, err
	}

	// ユーザーを作成する
	err = models.CreateUser(&models.User{
		UserID:       uid,
		Name:         args.Name,
		Email:        args.Email,
		PasswordHash: "",
		ProvUID:      args.ProviderUserID,
		CreatedAt:    now,
	}, models.ProviderCode(args.ProviderCode))

	// エラー処理
	if err != nil {
		return "", err
	}

	// トークンを生成
	token, err := NewSession(SessionArgs{
		UserID:    uid,
		RemoteIP:  args.RemoteIP,
		UserAgent: args.UserAgent,
	})

	// エラー処理
	if err != nil {
		return "", err
	}

	return token, nil
}
