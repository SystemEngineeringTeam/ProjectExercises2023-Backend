package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sensing struct {
	state string
}

func GetSensing(azimuth string) string {

	url := "http://localhost:10080/sensing?azimuth=" + azimuth

	// リクエストを生成
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// リクエストを実行
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// レスポンスボディを取得
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// レスポンスボディを表示
	fmt.Println(string(body))
	fmt.Println(jsonParth(string(body)))

	return jsonParth(string(body))

}

func jsonParth(body string) string {
	// jsonパース
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		fmt.Println(err)
	}

	// jsonの中身を表示
	fmt.Println(data["state"])

	return data["state"].(string)

}
