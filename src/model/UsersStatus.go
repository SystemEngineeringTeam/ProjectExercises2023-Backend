package model

import "time"

type UsersStatus struct {
	UsersStatusId uint `json:"id" gorm:"primaryKey"`
	Azimuth string `json:"azimuth"`
	Time time.Time `json:"time"`
	Status string `json:"status"`
}