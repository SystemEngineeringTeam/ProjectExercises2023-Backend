package model

import (
	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/lib"
)

// MYSQLの接続情報
var db = lib.SqlConnect()

// テーブル作成
func CreateAllTable() {
	// db.AutoMigrate(&Building{})
}
