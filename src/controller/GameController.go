package controller

import (
	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/model"
	"github.com/gin-gonic/gin"
	"time"
)

// GameStart ゲームを開始の処理
func GameStart(c *gin.Context) {
	req := model.BoardSurface{
		//現在の時間を取得
		StartTime:  time.Now(),
		FinishTime: time.Now(),
	}

	// データベースにデータを挿入
	model.CreateBoardSurface(&req)

	// データを返す
	c.JSON(200, gin.H{
		"StartTime":  req.StartTime,
		"FinishTime": req.FinishTime,
	})
}

// GameFinish ゲーム終了の処理
func GameFinish(c *gin.Context) {
	//最新のBoardSurfaceを取得
	latestBoardSurface := model.GetLatestBoardSurface()
	latestBoardSurface.FinishTime = time.Now() //現在の時間を割り当てる

	// データベースにデータを挿入
	model.UpdateBoardSurface(&latestBoardSurface)

	// データを返す
	c.JSON(200, gin.H{
		"FinishTime": latestBoardSurface.FinishTime,
	})
}
