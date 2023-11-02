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

// CanGameStart ゲームが開始できる状態か判定する
func CanGameStart() bool {
	//最新のBoardSurfaceを取得
	latestBoardSurface := GetLatestBoardSurface()

	//開始時刻と終了時刻を比較
	//等しくない場合=>前回のゲームが終了している
	if !latestBoardSurface.StartTime.Equal(latestBoardSurface.FinishTime) {
		return true //ゲーム開始ができる状態
	} else {
		return false
	}
}

// CanGameFinish ゲームが終了できる状態か判定する
func CanGameFinish() bool {
	return !CanGameStart()
}
