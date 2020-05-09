package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Login struct {
	Username string `form:"username" json:"user" binding:"required"`
	Password string `form:"password" json:"pwd" binding:"required"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("../template/**/*")
	r.Static("/static", "../img") // 参数1：URL代码中的文件前缀；参数：实际存放文件路径

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "book.html", gin.H{
			"message": "用户登录",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		login := Login{}
		err := c.ShouldBind(&login)
		if err != nil {
			log.Printf("%v\n", err)
			c.HTML(http.StatusNotFound, "book.html", gin.H{
				"message": err,
			})
			return
		}
		log.Printf("login success: %#v\n", login)
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": fmt.Sprintf("%#v", login),
		})
	})

	r.Run(":9091")

}
