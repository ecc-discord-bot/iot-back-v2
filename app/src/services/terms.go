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
	user, err := models.GetUser(claim.UserID)

	// エラー処理
	if err != nil {
		return err
	}

	// 同意済みか判定する
	if user.IsAgreed {
		return errors.New("terms already accepted")
	}

	// 同意済みにする
	user.IsAgreed = true
	user.Name = args.UserName
	user.Class = args.Class
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
		err = spreadsheet.WriteUser(fmt.Sprintf("管理シート!B%s", strconv.Itoa(result.Index)), spreadsheet.User{
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
		err = spreadsheet.WriteUser(fmt.Sprintf("管理シート!B%s", strconv.Itoa(result.Total)), spreadsheet.User{
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