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
	HeartRateDates []HeartRateData `gorm:"foreignKey:HeartRateId"`
	UsersStatuses  []UsersStatus   `gorm:"foreignKey:UsersStatusId"`
}

func CreateBoardSurface(boardSurface *BoardSurface) {
	result := db.Create(&boardSurface)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("boardSurface created!!", boardSurface)
}
