package model

import (
	"fmt"
	"log"
	"time"
)

type HeartRateData struct {
	HeartRateId    uint      `json:"id" gorm:"primaryKey"`
	BoardSurfaceId uint      `json:"board_surface_id"`
	Time           time.Time `json:"time"`
	Azimuth        string    `json:"azimuth"`
	HeartRate      int       `json:"bpm"`
}

// CreateHeartRateData DB上に新規作成
func CreateHeartRateData(heartRateData *HeartRateData) {
	result := db.Create(&heartRateData)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("heartRateData created!!", heartRateData)
}

// GetHeartRateData 最新のHeartRateDataを取得する
func GetHeartRateData(azimuth string) HeartRateData {
	// 最新のHeartRateDataを取得する
	targetHeartRateData := HeartRateData{}
	//db.Last(&targetHeartRateData) //代入

	//最新のBoardSurfaceIdを取得
	boardId := GetLastBoardId()

	//BoardSurfaceIdとazimuthが一致するものを取得
	db.Where("board_surface_id = ? AND azimuth = ?", boardId, azimuth).Last(&targetHeartRateData)
	//返却
	return targetHeartRateData
}
