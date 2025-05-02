package services

import (
	"app/logger"
	"app/middlewares"
	"app/models"
	"app/spreadsheet"
	"app/utils"
	"errors"
	"fmt"
	"strconv"
)

type AccepterTerms struct {
	UserName string `json:"UserName"`
	Class    string `json:"UserClass"`
}

func AcceptTerms(claim middlewares.AccessTokenClaim,args AccepterTerms) error {
	// ユーザーを取得する
	user, err := models.GetFromDiscordId(claim.ProvUid)

	// エラー処理
	if err != nil {
		return err
	}

	// 同意済みか判定する
	if user.IsAgreed {
		return errors.New("terms already accepted")
	}

	// ユーザーIDを更新
	user.UserID = claim.UserID

	// 同意済みにする
	user.IsAgreed = true

	// 名前を更新
	user.Name = args.UserName

	// クラスを更新
	user.Class = args.Class

	// 署名を更新
	user.Signature = args.UserName

	// ユーザーを更新する
	err = models.UpdateUser(user)

	// エラー処理
	if err != nil {
		return err
	}

	// spreadsheet から取得
	result, err := spreadsheet.GetLastRow(user.DiscordId)

	// エラー処理
	if err != nil {
		return err
	}

	// 見つかった時
	if result.Isfind {
		logger.Println("既存ユーザー")

		// spreadsheet に書き込む
		err = spreadsheet.WriteUser(fmt.Sprintf("%s!A%s",spreadsheet.BaseSheet, strconv.Itoa(result.Index)), spreadsheet.User{
			UserID:     user.UserID,
			DiscordID:  user.DiscordId,
			StudentsID: user.StudentsId,
			Name:       args.UserName,
			Class:      args.Class,
			IsPaid:     false,
			IsAgreed:   true,
			Time:       utils.NowTime(),
			Signature:  args.UserName,
		})

		// エラー処理
		if err != nil {
			return err
		}
	} else {
		logger.Println("新規ユーザー")
		// spreadsheet に書き込む
		err = spreadsheet.WriteUser(fmt.Sprintf("%s!A%s",spreadsheet.BaseSheet, strconv.Itoa(result.Total)), spreadsheet.User{
			UserID:     user.UserID,
			DiscordID:  user.DiscordId,
			StudentsID: user.StudentsId,
			Name:       args.UserName,
			Class:      args.Class,
			IsPaid:     false,
			IsAgreed:   true,
			Time:       utils.NowTime(),
			Signature:  args.UserName,
		})

		// エラー処理
		if err != nil {
			return err
		}
	}

	return nil
}