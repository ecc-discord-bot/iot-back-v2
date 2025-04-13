package services

import (
	"app/middlewares"
	"app/models"
	"app/utils"
	"errors"
	"strings"

	"github.com/markbates/goth"
)

func LinkAccount(claim middlewares.AccessTokenClaim,gothUser goth.User) error {

	// ユーザーを取得する
	_, err := models.GetUser(claim.UserID)

	// エラー処理
	if err == nil {
		// ユーザーが存在する場合
		return errors.New("user already linked")
	}

	// メールアドレスから @ecc.ac.jp を除去する
	studentId := gothUser.Email[:strings.Index(gothUser.Email, "@ecc.ac.jp")]

	// 存在ない場合
	user := models.User{
		UserID:     claim.UserID,
		DiscordId:  claim.ProvUid,
		StudentsId: studentId,
		Name:       "",
		Class:      "",
		Signature:  "",
		NowTime:    utils.NowTime(),
		IsPaid:     false,
		IsAgreed:   false,
	}

	// ユーザーを作成する
	return models.CreateUser(user)
}
