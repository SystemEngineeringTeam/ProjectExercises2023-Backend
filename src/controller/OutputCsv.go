package controller

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/model"
)

func AllBeartOutput() {

	boardId := int(model.GetLastBoardId())

	BeartOutput("north", boardId)
	BeartOutput("south", boardId)
	BeartOutput("east", boardId)
	BeartOutput("west", boardId)
}

func AllUserStatusOutput() {

	boardId := int(model.GetLastBoardId())

	UserStatusOutput("north", boardId)
	UserStatusOutput("south", boardId)
	UserStatusOutput("east", boardId)
	UserStatusOutput("west", boardId)
}

func BeartOutput(azimuth string, boardId int) {

	//ディレクトリの作成
	path := "../output_csv/" + strconv.Itoa(boardId) + "/" + azimuth + "_bpm" + ".csv"

	//mkdir "
	CreateDirectory("../output_csv/" + strconv.Itoa(boardId))

	// データを取得
	heartRateList := model.GetAllHeartRateData(azimuth)

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// csvに出力
	w := csv.NewWriter(file)

	// データを書き込む
	w.Write([]string{"time", "bpm"})
	for _, v := range heartRateList {
		w.Write([]string{strconv.FormatInt(v.Time.Unix(),10), strconv.Itoa(v.HeartRate)})
	}
	w.Flush()
	fmt.Println("CSV出力完了")
}

func UserStatusOutput(azimuth string, boardId int) {
	//ディレクトリの作成
	path := "../output_csv/" + strconv.Itoa(boardId) + "/" + azimuth + "_emotion" + ".csv"

	//mkdir "
	CreateDirectory("../output_csv/" + strconv.Itoa(boardId))

	// データを取得
	userStatusList := model.GetAllUsersStatus(azimuth)

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// csvに出力
	w := csv.NewWriter(file)

	// データを書き込む
	w.Write([]string{"time", "status"})
	for _, v := range userStatusList {
		w.Write([]string{strconv.FormatInt(v.Time.Unix(),10), v.Status})
	}
	w.Flush()
	fmt.Println("CSV出力完了")
}

// ディレクトリを作成する
func CreateDirectory(path string) {
	fileInfo, err := os.Lstat("./")
	if err != nil {
		log.Fatal(err)
	}
	fileMode := fileInfo.Mode()
	nixPerms := fileMode & os.ModePerm
	if err := os.MkdirAll(path, nixPerms); err != nil {
		log.Fatal(err)
	}
}
