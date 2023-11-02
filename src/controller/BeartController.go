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

	//関数実行可能かどうかを判定
	if !IsExecutable(c, targetAzimuth) {
		return
	}

	//URLの方位をチェック
	if !checkAzimuth(targetAzimuth) {
		c.JSON(400, gin.H{
			"message": "不正な方位です",
		})
		return
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

	//関数実行可能かどうかを判定
	if !IsExecutable(c, targetAzimuth) {
		return
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

	//関数実行可能かどうかを判定
	if !IsExecutable(c, targetAzimuth) {
		return
	}

	//感情
	emotions := [...]string{
		"normal",   //平常
		"surprise", //驚愕
		"nervous",  //緊張
		"relief",   //安堵
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

	//関数実行可能かどうかを判定
	if !IsExecutable(c, targetAzimuth) {
		return
	}

	//最新のBoardSurfaceを取得
	latestUsersStatus := model.GetUsersStatus(targetAzimuth)

	// データを返す
	c.JSON(200, latestUsersStatus)
}

// URLから取り出した方位に不正な文字列が入っていないかチェックする関数
func checkAzimuth(targetAzimuth string) bool {
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
		//最後まで一致しなかったときfalse(不正)を返す
		if i == len(azimuth)-1 {
			return false
		}
	}

	//最後まで一致したときtrue(正常)を返す
	return true
}

// IsExecutable 関数が実行可能かどうかを判定する関数
func IsExecutable(c *gin.Context, targetAzimuth string) bool {
	//ゲームが継続しているかチェック
	if !model.IsGameContinuing() {
		c.JSON(400, gin.H{
			"message": "ゲームはまだ開始されていません",
		})
		return false
	}

	//URLの方位をチェック
	if !checkAzimuth(targetAzimuth) {
		c.JSON(400, gin.H{
			"message": "不正な方位です",
		})
		return false
	}

	//関数実行可能
	return true
}
