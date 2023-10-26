package router

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	// サーバーログの出力先を設定
	gin.DisableConsoleColor()
	f, _ := os.Create("../server.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// ルーティングの設定
	router := gin.Default()

	// ここからCorsの設定
	router.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	// V1の設定
	v1 := router.Group("/api/v1/")

	//ゲーム開始
	v1.GET("/start", controller.GameStart)
	//ゲーム終了
	v1.GET("/finish", controller.GameFinish)
	//心拍数の送信
	azimuths := [...]string{
		"north", //北
		"south", //南
		"east",  //東
		"west",  //西
	}
	for _, azimuth := range azimuths {
		v1.GET(fmt.Sprintf("set_bpm/%s", azimuth), controller.SendHeartRate)
	}

	// 下記を追記することで`http://localhost:8080/api/v1/swagger/index.html`を叩くことでswagger uiを開くことができる
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// ポート番号の設定
	router.Run(":8080")
}
