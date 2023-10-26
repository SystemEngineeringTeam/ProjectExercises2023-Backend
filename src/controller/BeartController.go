package controller

import (
	"fmt"
	"time"

	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/model"
	"github.com/gin-gonic/gin"
)

// GameStart ゲームを開始する。
func GameStart(c *gin.Context) {
	fmt.Println("gamestart")
	req := model.BoardSurface{
		//現在の時間を取得
		StartTime:  time.Now(),
		FinishTime: time.Now(),
	}

	// データベースにデータを挿入
	model.CreateBoardSurface(&req)

	// データを返す
	c.JSON(200, gin.H{
		"time": req.StartTime,
	})
}

// GameFinish ゲーム終了の処理
func GameFinish(c *gin.Context) {
	fmt.Println("gamefinish")
	req := model.BoardSurface{
		//現在の時間を取得
		FinishTime: time.Now(),
	}

	//最新のBoardSurfaceを取得
	latestBoardSurface := model.GetLatestBoardSurface()
	latestBoardSurface.FinishTime = time.Now() //現在の時間を割り当てる

	// データベースにデータを挿入
	model.UpdateBoardSurface(&latestBoardSurface)

	// データを返す
	c.JSON(200, gin.H{
		"finishTime": req.FinishTime,
	})
}
