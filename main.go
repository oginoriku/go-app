package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 自動的にファイルを返すよう設定 --- (*1)
	router.StaticFS("/view", http.Dir("view"))

	// ルートなら /static/index.html にリダイレクト --- (*2)
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "/view/index.html")
	})

	// フォームの内容を受け取って挨拶する --- (*3)
	router.GET("/hello", func(ctx *gin.Context) {
		name := ctx.Query("name")
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>Hello, "+name+"</h1>")
	})

	// サーバーを起動
	err := router.Run("127.0.0.1:8888")
	if err != nil {
		log.Fatal("サーバー起動に失敗", err)
	}
}
