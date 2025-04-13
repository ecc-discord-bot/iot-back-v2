package services

import "auth/models"

func GetAccessToken(userID string) (string, error) {
	// ユーザーを取得
	user, err := models.GetUser(userID)

	// エラー処理
	if err != nil {
		return "", err
	}

	// トークンを生成
	token, err := AccessTokenJwt(AccessTokenClaim{UserID: userID, Labels: []string{}, ProvCode: user.ProvCode, ProvUid: user.ProvUID})

	return token, err
}
