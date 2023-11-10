package model

import (
	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/src/lib"
)

/**
テーブル作成するやつ
*/

// MYSQLの接続情報
var db = lib.SqlConnect()

// テーブル作成
func CreateAllTable() {
	db.AutoMigrate(&BoardSurface{})
	db.AutoMigrate(&HeartRateData{})
	db.AutoMigrate(&UsersStatus{})
}
