package main

import (
	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/src/model"
	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/src/router"
)

// @title gin-swagger todos
// @version 1.0
// @license.name kosuke
// @description このswaggerはgin-swaggerの見本apiです
func main() {
	// controller.AddCsvData()

	// テーブル作成とDB接続
	model.CreateAllTable()

	// ルーティングの設定＋サーバー起動
	router.Init()

}
