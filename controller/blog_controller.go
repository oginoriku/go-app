package controller

import (
	"github.com/gin-gonic/gin"
)

func IndexFunc(c *gin.Context) {
	c.Redirect(302, "/view/index.html")
}

func HelloFunc(c *gin.Context) {
	name := c.Query("name")
	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.String(200, "<h1>Hello, "+name+"</h1>")
}
