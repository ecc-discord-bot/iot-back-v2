package services

import (
	"app/models"
	"app/spreadsheet"

	"gorm.io/gorm"
)

func SyncData() error {
	// データを取得
	datas, err := spreadsheet.SyncDatas()

	// エラー処理
	if err != nil {
		return err
	}

	for _, data := range datas {
		// ユーザーを作成する
		user := models.User{
			UserID:     data.UserID,
			DiscordId:  data.DiscordID,
			StudentsId: data.StudentID,
			Name:       data.Name,
			Class:      data.Class,
			Signature:  data.Signature,
			NowTime:    data.AgreeTime,
			IsPaid:     data.IsPaid,
			IsAgreed:   data.IsAgreed,
		}

		// ユーザーを取得
		_, err := models.GetFromDiscordId(data.DiscordID)

		// レコードが存在しない場合
		if err == gorm.ErrRecordNotFound {
			// ユーザーを作成する
			if err := models.CreateUser(user); err != nil {
				return err
			}

			continue
		}

		// エラー処理
		if err != nil {
			return err
		}

		// ユーザーを更新する
		if err := models.UpdateUser(user); err != nil {
			return err
		}
	}

	return nil
}