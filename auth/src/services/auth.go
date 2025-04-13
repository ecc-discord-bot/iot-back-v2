package services

import "auth/models"

func Logout(session *models.Session) error {
	// ユーザー取得
	user, err := models.GetUser(session.UserID)

	// エラー処理
	if err != nil {
		return err
	}

	// セッションを削除
	return user.DeleteSession(session.SessionID)
}
