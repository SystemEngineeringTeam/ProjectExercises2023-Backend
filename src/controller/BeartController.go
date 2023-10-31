package controller

import (
	"fmt"
	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/model"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

// SendHeartRate 心拍数を送る
func SendHeartRate(c *gin.Context) {
	//URLから方位を取り出す
	targetAzimuth := c.Param("azimuth")
	//各方位
	azimuths := [...]string{
		"north", //北
		"south", //南
		"east",  //東
		"west",  //西
	}

	//4方位以外の文字例が入っていないかチェック
	for i, azimuth := range azimuths {
		//一致するときfor文を抜ける
		if targetAzimuth == azimuth {
			break
		}
		//最後まで一致しなかったとき関数を終了する
		if i == len(azimuth)-1 {
			c.JSON(500, gin.H{
				"type":    "failed",
				"message": "東南西北がありません。",
			})
			return
		}
	}

	req := model.HeartRateData{
		BoardSurfaceId: model.GetLastBoardId(),
		Time:           time.Now(),
		Azimuth:        targetAzimuth,
		HeartRate:      rand.Intn(100), //適当
	}

	// データベースにデータを挿入
	model.CreateHeartRateData(&req)

	// データを返す
	c.JSON(200, gin.H{
		"BoardSurfaceId": req.BoardSurfaceId,
		"time":           req.Time,
		"Azimuth":        req.Azimuth,
		"HeartRate":      req.HeartRateId,
	})
}

// GetHeartRate 心拍数を取得する
func GetHeartRate(c *gin.Context) {
	//URLから方位を取り出す
	targetAzimuth := c.Param("azimuth")
	//各方位
	azimuths := [...]string{
		"north", //北
		"south", //南
		"east",  //東
		"west",  //西
	}

	//4方位以外の文字例が入っていないかチェック
	for i, azimuth := range azimuths {
		//一致するときfor文を抜ける
		if targetAzimuth == azimuth {
			break
		}
		//最後まで一致しなかったとき関数を終了する
		if i == len(azimuth)-1 {
			return
		}
	}

	//最新のBoardSurfaceを取得
	latestHeartRateData := model.GetHeartRateData(targetAzimuth)

	fmt.Println(latestHeartRateData)
	// データを返す
	c.JSON(200, latestHeartRateData)
}

// SendEmotionStatus 心情を送る
func SendEmotionStatus(c *gin.Context) {
	//URLから方位を取り出す
	targetAzimuth := c.Param("azimuth")
	//各方位
	azimuths := [...]string{
		"north", //北
		"south", //南
		"east",  //東
		"west",  //西
	}
	//感情
	emotions := [...]string{
		"normal",   //平常
		"surprise", //驚愕
		"nervous",  //緊張
		"relief",   //安堵
	}
	//4方位以外の文字例が入っていないかチェック
	for i, azimuth := range azimuths {
		//一致するときfor文を抜ける
		if targetAzimuth == azimuth {
			break
		}
		//最後まで一致しなかったとき関数を終了する
		if i == len(azimuth)-1 {
			c.JSON(500, gin.H{
				"type":    "failed",
				"message": "東南西北がありません。",
			})
			return
		}
	}

	req := model.UsersStatus{
		BoardSurfaceId: model.GetLastBoardId(),
		Time:           time.Now(),
		Azimuth:        targetAzimuth,
		Status:         emotions[rand.Intn(4)], //適当
	}

	// データベースにデータを挿入
	model.CreateUsersStatus(&req)

	// データを返す
	c.JSON(200, gin.H{
		"BoardSurfaceId": req.BoardSurfaceId,
		"time":           req.Time,
		"Azimuth":        req.Azimuth,
		"Status":         req.Status,
	})
}

// GetEmotionStatus	最新の心情を取得する
func GetEmotionStatus(c *gin.Context) {
	//URLから方位を取り出す
	targetAzimuth := c.Param("azimuth")
	//各方位
	azimuths := [...]string{
		"north", //北
		"south", //南
		"east",  //東
		"west",  //西
	}

	//4方位以外の文字例が入っていないかチェック
	for i, azimuth := range azimuths {
		//一致するときfor文を抜ける
		if targetAzimuth == azimuth {
			break
		}
		//最後まで一致しなかったとき関数を終了する
		if i == len(azimuth)-1 {
			return
		}
	}

	fmt.Println(targetAzimuth)

	//最新のBoardSurfaceを取得
	latestUsersStatus := model.GetUsersStatus(targetAzimuth)

	// データを返す
	c.JSON(200, latestUsersStatus)
}
