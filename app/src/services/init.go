package services

import (
	"app/spreadsheet"
)

func Init() {
	// スプレッドシート初期化
	spreadsheet.SpreadsheetInit()

	// マイグレーション
	// spreadsheet.Migrate()
}
