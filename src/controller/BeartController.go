package controller

import (
	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/model"
	"github.com/gin-gonic/gin"
	"math/rand"
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
		"time": req.StartTime,
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
		"finishTime": latestBoardSurface.FinishTime,
	})
}

// SendHeartRate 心拍数を送る
func SendHeartRate(c *gin.Context) {
	//各方位
	azimuths := [...]string{
		"north", //北
		"south", //南
		"east",  //東
		"west",  //西
	}
	req := model.HeartRateData{
		Time:      time.Now(),
		Azimuth:   azimuths[rand.Intn(4)],
		HeartRate: rand.Intn(100),
	}

	// データベースにデータを挿入
	model.CreateHeartRateData(&req)

	// データを返す
	c.JSON(200, gin.H{
		"time":      req.Time,
		"Azimuth":   req.Azimuth,
		"HeartRate": req.HeartRateId,
	})
}
