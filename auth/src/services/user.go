package services

import "auth/models"

type UserInfo struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ProvCode string `json:"prov_code"`
	ProvUid  string `json:"prov_uid"`
}

func GetMe(userid string) (UserInfo, error) {
	// ユーザー取得
	user, err := models.GetUser(userid)

	// エラー処理
	if err != nil {
		return UserInfo{}, err
	}

	return UserInfo{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		ProvCode: string(user.ProvCode),
		ProvUid:  user.ProvUID,
	}, nil
}
