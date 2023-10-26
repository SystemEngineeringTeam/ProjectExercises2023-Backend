package model

import (
	"fmt"
	"log"
	"time"
)

type HeartRateData struct {
	HeartRateId uint      `json:"id" gorm:"primaryKey"`
	Time        time.Time `json:"time"`
	Azimuth     string    `json:"azimuth"`
	HeartRate   int       `json:"heart_rate"`
}

// CreateHeartRateData DB上に新規作成
func CreateHeartRateData(heartRateData *HeartRateData) {
	result := db.Create(&heartRateData)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("heartRateData created!!", heartRateData)
}
