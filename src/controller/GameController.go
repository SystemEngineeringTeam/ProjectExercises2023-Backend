package controller

import (
	"strconv"
	"time"

	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/src/model"
	"github.com/gin-gonic/gin"
)

// GameStart ゲーム開始の処理
// @Summary Todo一覧を配列で返す
// @Produce  json
// @Success 200 {object} model.BoardSurface
// @Router /todos [get]
func GameStart(c *gin.Context) {
	//ゲームが継続している場合
	if model.IsGameContinuing() {
		c.JSON(400, gin.H{
			"message": "ゲームはすでに開始されています",
		})
		return
	}

	req := model.BoardSurface{
		//現在の時間を取得
		StartTime:  time.Now(),
		FinishTime: time.Now(),
	}

	// データベースにデータを挿入
	model.CreateBoardSurface(&req)

	// データを返す
	c.JSON(200, gin.H{
		"start_time":  req.StartTime,
		"finish_time": req.FinishTime,
	})
}

// GameFinish ゲーム終了の処理
func GameFinish(c *gin.Context) {
	//ゲームが継続していない場合
	if !model.IsGameContinuing() {
		c.JSON(400, gin.H{
			"message": "ゲームはまだ開始されていません",
		})
		return
	}
	//最新のBoardSurfaceを取得
	latestBoardSurface := model.GetLatestBoardSurface()
	latestBoardSurface.FinishTime = time.Now() //現在の時間を割り当てる

	// データベースにデータを挿入
	model.UpdateBoardSurface(&latestBoardSurface)

	// データをCSVに出力
	AllBeartOutput()
	AllUserStatusOutput()

	GenerateGraph(strconv.Itoa(int(model.GetLastBoardId())))

	// データを返す
	c.JSON(200, gin.H{
		"finish_time": latestBoardSurface.FinishTime,
		"url":         "http://heartbeat.sysken.net/api/v1/get/image/" + strconv.Itoa(int(model.GetLastBoardId())),
	})
}
