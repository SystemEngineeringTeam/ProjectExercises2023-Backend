package model

import (
	"fmt"
	"log"
	"time"
)

type UsersStatus struct {
	UsersStatusId  uint      `json:"id" gorm:"primaryKey"`
	BoardSurfaceId uint      `json:"board_surface_id"`
	Azimuth        string    `json:"azimuth"`
	Time           time.Time `json:"time"`
	Status         string    `json:"status"`
}

// CreateUsersStatus DB上に新規作成
func CreateUsersStatus(usersStatus *UsersStatus) {
	result := db.Create(&usersStatus)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("usersStatus created!!", usersStatus)
}

// GetUsersStatus 最新のHeartRateDataを取得する
func GetUsersStatus(azimuth string) UsersStatus {
	// 最新のHeartRateDataを取得する
	targetUsersStatus := UsersStatus{}
	//db.Last(&targetUsersStatus) //代入

	//最新のBoardSurfaceIdを取得
	boardId := GetLastBoardId()

	//BoardSurfaceIdとazimuthが一致するものを取得
	db.Where("board_surface_id = ? AND azimuth = ?", boardId, azimuth).Last(&targetUsersStatus)
	//返却
	return targetUsersStatus
}

func GetAllUsersStatus(azimuth string) []UsersStatus {
	// 最新のHeartRateDataを取得する
	var targetUsersStatus []UsersStatus

	//最新のBoardSurfaceIdを取得
	boardId := GetLastBoardId()

	//BoardSurfaceIdとazimuthが一致するものを取得
	db.Where("board_surface_id = ? AND azimuth = ?", boardId, azimuth).Find(&targetUsersStatus)
	//返却
	return targetUsersStatus
}