package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HandlerImage(ctx *gin.Context) {
	imageID := ctx.Param("imageID")

	//パスをここで指定
	path := "output_image/" + imageID + ".png"

	fmt.Println(path)

	//ここで画像を返す
	ctx.File(path)
}
