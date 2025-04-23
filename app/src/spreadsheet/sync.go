package spreadsheet

import (
	"app/logger"
	"app/models"
	"fmt"
	"strconv"
	"time"
)

// クラス
var classes = []string{}

func SyncClass() error {
	//値取得
	resp, err := service.Spreadsheets.Values.Get(spreadsheetID, "設定!E3:E").Do()

	//エラー処理
	if err != nil {
		return err
	}

	//初期化
	classes = []string{}

	//ループする
	for _, row := range resp.Values {
		//0この時
		if len(row) == 0 {
			continue
		}

		//値を追加
		classes = append(classes, row[0].(string))
	}

	return nil
}

type SheetData struct {
	UserID       string //ユーザーID
	DiscordID    string //DiscordID
	StudentID    string //学籍番号
	Name         string //名前
	Class        string //クラス
	IsAgreed     bool   //同意
	IsPaid       bool   //支払い
	AgreeTime    int64  //同意時間
	Signature    string //署名
	IsRetirement bool   //退会フラグ
}

func SyncDatas() ([]SheetData, error) {
	// 最終行を取得する
	result, err := GetLastRow("")

	// エラー処理
	if err != nil {
		return []SheetData{}, err
	}

	// 値を取得する
	values, err := service.Spreadsheets.Values.Get(spreadsheetID, fmt.Sprintf("管理シート!A3:K%s", strconv.Itoa(result.Total-1))).DateTimeRenderOption("FORMATTED_STRING").Do()

	// エラー処理
	if err != nil {
		return []SheetData{}, err
	}

	datas := []SheetData{}

	// 値を取得する
	for _, row := range values.Values {
		// 0この時
		if len(row) == 0 {
			continue
		}

		logger.Println(row)

		// 値を追加
		datas = append(datas, SheetData{
			UserID:       row[0].(string),
			DiscordID:    row[1].(string),
			StudentID:    row[2].(string),
			Name:         row[3].(string),
			Class:        row[4].(string),
			IsAgreed:     row[5].(string) == "TRUE",
			IsPaid:       row[6].(string) == "TRUE",
			AgreeTime:    ConvertTime(row[7].(string)),
			Signature:    row[8].(string),
			IsRetirement: row[9].(string) == "TRUE",
		})
	}

	logger.Println(datas)

	return datas, nil
}

func ConvertTime(timeStr string) int64 {
	// タイムスタンプのフォーマット定義
	layout := "2006/01/02 15:04:05"

	// time.Parse を使って文字列からtime.Time型に変換
	timeData, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Printf("エラーが発生しました: %v\n", err)
		return 0
	}

	// time.Time を int64 に変換
	timeInt := timeData.Unix()
	return timeInt
}

func Migrate() error {
	// 全ユーザー取得
	users, err := models.GetAllUsers()
	if err != nil {
		return err
	}

	// ユーザーを回す
	for _, user := range users {
		err = postUser(user)

		// エラー処理
		if err != nil {
			logger.Println(err)
			continue
		}
	}

	return nil
}

func postUser(user models.User) error {
	// spreadsheet から取得
	result, err := GetLastRow(user.DiscordId)

	// エラー処理
	if err != nil {
		return err
	}

	// 見つかった時
	if result.Isfind {
		logger.Println("既存ユーザー")

		// spreadsheet に書き込む
		err = WriteUser(fmt.Sprintf("管理シート!A%s", strconv.Itoa(result.Index)), User{
			UserID:     user.UserID,
			DiscordID:  user.DiscordId,
			StudentsID: user.StudentsId,
			Name:       user.Name,
			Class:      user.Class,
			IsPaid:     false,
			IsAgreed:   true,
			Time:       user.NowTime,
			Signature:  user.Signature,
		})

		// エラー処理
		if err != nil {
			return err
		}
	} else {
		logger.Println("新規ユーザー")
		// spreadsheet に書き込む
		err = WriteUser(fmt.Sprintf("管理シート!A%s", strconv.Itoa(result.Total)), User{
			UserID:     user.UserID,
			DiscordID:  user.DiscordId,
			StudentsID: user.StudentsId,
			Name:       user.Name,
			Class:      user.Class,
			IsPaid:     user.IsPaid,
			IsAgreed:   user.IsAgreed,
			Time:       user.NowTime,
			Signature:  user.Signature,
		})

		// エラー処理
		if err != nil {
			return err
		}
	}

	return nil
}