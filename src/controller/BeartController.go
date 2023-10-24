package controller

import (
	"time"

	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/model"
	"github.com/gin-gonic/gin"
)

func GameStart(c *gin.Context) {

	req := model.BoardSurface{
		//現在の時間を取得
		StartTime: time.Now(),
		FinishTime: time.Now(),
	}

	// データベースにデータを挿入
	model.CreateBoardSurface(&req)

	// データを返す
	c.JSON(200, gin.H{
		"time" : req.StartTime,
	})
}
