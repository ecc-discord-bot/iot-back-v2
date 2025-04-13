package services

import (
	"os"
)

var (
	// トークンシークレット
	TokenSecret = "secret"
)

func Init() {
	// 環境変数からトークンシークレットを取得
	TokenSecret = os.Getenv("TOKEN_SECRET")

	// 環境変数から秘密鍵を取得
	certString := os.Getenv("JWT_PRIVATE_KEY")

	// 秘密鍵を初期化
	initJwt(certString)
}