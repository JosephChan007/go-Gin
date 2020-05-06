package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexAction(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "index page...",
	})
}

func main() {

	router := gin.Default()

	// 装载HTML文件
	router.LoadHTMLGlob("../template/**/*")
	router.Static("/static", "../img") // 参数1：URL代码中的文件前缀；参数：实际存放文件路径

	router.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hello.html", gin.H{
			"message": "Hello world!!",
		})
	})

	router.GET("/index", indexAction)

	router.Run(":9091")
}
