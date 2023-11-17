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

	//controller.AddCsvData("./output_csv/23/east_bpm.csv")
	//controller.AddCsvData("./output_csv/23/west_bpm.csv")
	//controller.AddCsvData("./output_csv/23/south_bpm.csv")
	//controller.AddCsvData("./output_csv/23/north_bpm.csv")

	// ルーティングの設定＋サーバー起動
	router.Init()

}
