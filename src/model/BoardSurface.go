package model

import (
	"time"

	"gorm.io/gorm"
)

type BoardSurface struct {
	gorm.Model
	Id uint `json:"id" gorm:"primaryKey"`
	Azimuth   string `json:"azimuth"`
	StartTime time.Time `json:"start_time"`
	FinishTime   time.Time `json:"finish_time"`
	HeartRateDates []HeartRateData `gorm:"foreignKey:HeartRateId"`
	UsersStatuses []UsersStatus `gorm:"foreignKey:UsersStatusId"`
}
