package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	//r.LoadHTMLGlob("view/*html")
	r.StaticFS("/view", http.Dir("view"))
	r.Static("/css", "resources/css/")
	r.GET("/", IndexFunc)
	r.GET("/hello", HelloFunc)
	return r
}
