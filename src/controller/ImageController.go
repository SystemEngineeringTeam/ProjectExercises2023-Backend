package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"time"
)

func HandlerImage(ctx *gin.Context) {
	imageID := ctx.Param("imageID")

	//パスをここで指定
	path := "output_image/" + imageID + ".png"

	fmt.Println(path)

	//ここで画像を返す
	ctx.File(path)
}

func GenerateGraph(boardId string) {
	fmt.Println(boardId)

	// コンテキストを作成し、キャンセル関数を取得
	ctx, cancel := context.WithCancel(context.Background())

	// コンテキストを cmd に設定
	out, err2 := exec.Command("pwd").Output()
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(out))
	cmd := exec.CommandContext(ctx, "python", "python_src/main.py", boardId)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// キャンセル関数を cmd.WaitDelay で設定
	cmd.WaitDelay = 5 * time.Second
	go func() {
		time.Sleep(5 * time.Second) // もしキャンセルしない場合、5秒後にキャンセルを呼び出す
		cancel()
	}()

	// 実行
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	// cmd.Wait()でプロセスが終了するまで待つ
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}

	// 処理が終わった後に1秒Sleep
	time.Sleep(1 * time.Second)
}
