package model

import "time"

type HeartRateData struct {
	HeartRateId uint `json:"id" gorm:"primaryKey"`
	Time time.Time `json:"time"`
	Azimuth string `json:"azimuth"`
	HeartRate int `json:"heart_rate"`
}