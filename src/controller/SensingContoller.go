package controller

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/src/model"
	"github.com/go-gota/gota/dataframe"
)

func CheckEmotion(list []model.HeartRateData) string {
	bpmList := convertBpmList(list)
	diff := diff(bpmList)

	fmt.Println("bpmList（差分）")
	fmt.Println(diff)

	// ここから感情判定
	//"normal",   //平常 1未満,-1より大きい
	// "surprise", //驚愕  1以上,
	// "relief",   //安堵 -1以下,

	if len(list) <= 1 {
		return "normal"
	}

	if diff[len(diff)-1].Bpm < 1 && diff[len(diff)-1].Bpm > -1 {
		return "normal"
	} else if diff[len(diff)-1].Bpm >= 1 {
		return "surprise"
	} else if diff[len(diff)-1].Bpm <= -1 {
		return "relief"
	} else {
		return "normal"
	}
}

func convertBpmList(list []model.HeartRateData) []model.BpmListModel {
	var bpmList []model.BpmListModel
	for _, v := range list {
		bpmList = append(bpmList, model.BpmListModel{
			Bpm:  float32(v.HeartRate),
			Time: v.Time,
		})
	}
	return bpmList
}

func diff(list []model.BpmListModel) []model.BpmListModel {
	var diffList []model.BpmListModel
	for i := 0; i < len(list)-1; i++ {
		diffList = append(diffList, model.BpmListModel{
			Bpm:  (list[i+1].Bpm - list[i].Bpm) / float32(list[i+1].Time.Sub(list[i].Time).Seconds()),
			Time: list[i+1].Time,
		})
	}
	return diffList
}

func AddCsvData() {
	// Read file
	f, err := os.Open("../test_csv/HartRate.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// File to Dataframe
	df := dataframe.ReadCSV(f)
	fmt.Println("dfをそのまま")
	fmt.Println(df)
	fmt.Println("df.Describe()")
	fmt.Println(df.Describe())

	// time, bpm という形でデータが入っているのでHeartRateData型に変換する
	// その後、DBに保存する
	var list []model.HeartRateData

	for i := 0; i < df.Nrow(); i++ {
		// 行を取得
		bpm := int(df.Subset(i).Elem(0, 1).Float())
		fmt.Println(bpm)

		time := int64(df.Subset(i).Elem(0, 0).Float() / 1000)
		fmt.Println(time)

		date := unixTime2Time(time)

		list = append(list, model.HeartRateData{
			BoardSurfaceId: model.GetLastBoardId(),
			Time:           date,
			Azimuth:        "north",
			HeartRate:      bpm,
		})
	}

	// DBに保存
	model.CreateAllHeartRateData(list)
}

func unixTime2Time(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}
