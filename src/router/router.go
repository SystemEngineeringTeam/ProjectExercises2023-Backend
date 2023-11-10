package router

import (
	"io"
	"os"
	"time"

	"github.com/SystemEngineeringTeam/ProjectExercises2023-Backend/src/controller"
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
	v1.POST("/start", controller.GameStart)
	//ゲーム終了
	v1.POST("/finish", controller.GameFinish)
	//TODO: データがないときに取得したらエラーを返すようにする
	//最新の心拍数の送信
	v1.POST("/set_bpm/:azimuth", controller.SendHeartRate)
	//最新の心拍数の取得
	v1.GET("/get_bpm/:azimuth", controller.GetHeartRate)
	//最新の心情の送信
	v1.POST("/set/user_status/:azimuth", controller.SendEmotionStatus)
	//最新の心情の取得
	v1.GET("/get/user_status/:azimuth", controller.GetEmotionStatus)
	//画像を返す
	v1.GET("/get/image/:imageID", controller.HandlerImage)

	// 下記を追記することで`http://localhost:8080/api/v1/swagger/index.html`を叩くことでswagger uiを開くことができる
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// ポート番号の設定
	router.Run(":8080")
}
