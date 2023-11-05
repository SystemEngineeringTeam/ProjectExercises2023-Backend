package model

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type BoardSurface struct {
	gorm.Model
	Id             uint            `json:"id" gorm:"primaryKey"`
	StartTime      time.Time       `json:"start_time"`
	FinishTime     time.Time       `json:"finish_time"`
	HeartRateDates []HeartRateData `gorm:"foreignKey:BoardSurfaceId"`
	UsersStatuses  []UsersStatus   `gorm:"foreignKey:BoardSurfaceId"`
}

func CreateBoardSurface(boardSurface *BoardSurface) {
	result := db.Create(&boardSurface)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("boardSurface created!!", boardSurface)
}

// GetLatestBoardSurface 最新のBoardSurfaceを取得する
func GetLatestBoardSurface() BoardSurface {
	// 最新のboardSurfaceを取得する
	targetBoardSurface := BoardSurface{}
	db.Last(&targetBoardSurface) //代入

	//返却
	return targetBoardSurface
}

func UpdateBoardSurface(boardSurface *BoardSurface) {
	//finishTimeを上書き
	result := db.Save(&boardSurface)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("boardSurface updated!!", boardSurface)
}

// GetLastBoardId 最後のBoardIdを取得する
func GetLastBoardId() uint {
	// 最新のboardSurfaceを取得する
	targetBoardSurface := BoardSurface{}
	db.Last(&targetBoardSurface) //代入

	//返却
	return targetBoardSurface.Id
}

// IsGameContinuing ゲームが継続しているか判定する
func IsGameContinuing() bool {
	//最新のBoardSurfaceを取得
	latestBoardSurface := GetLatestBoardSurface()

	//latestBoardSurfaceの取得に失敗した場合
	if latestBoardSurface.Id == 0 {
		return false //ゲームが継続していない
	}

	//開始時刻と終了時刻を比較
	//等しい場合=>ゲームが継続している
	if latestBoardSurface.StartTime.Equal(latestBoardSurface.FinishTime) {
		return true //ゲームが継続している
	} else {
		return false //ゲームが継続していない
	}
}
