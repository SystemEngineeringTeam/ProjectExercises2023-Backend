package model

import "time"

type UsersStatus struct {
	UsersStatusId  uint      `json:"id" gorm:"primaryKey"`
	BoardSurfaceId uint      `json:"board_surface_id"`
	Azimuth        string    `json:"azimuth"`
	Time           time.Time `json:"time"`
	Status         string    `json:"status"`
}
