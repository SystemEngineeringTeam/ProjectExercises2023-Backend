package main

import (
	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/model"
	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/router"
)

// @title CampusCrowdMonitor
// @version 1.0
// @description このswaggerはCampusCrowdMonitorのapiです
func main() {
	// テーブル作成とDB接続
	model.CreateAllTable()
	// ルーティングの設定＋サーバー起動
	router.Init()
}
