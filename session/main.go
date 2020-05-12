package main

import (
	mgr "github.com/JosephChan007/go-Gin/session/manager"
	mid "github.com/JosephChan007/go-Gin/session/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("../template/**/*")
	r.Static("/static", "../img") // 参数1：URL代码中的文件前缀；参数：实际存放文件路径

	r.GET("/404", func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "session/404.html", nil)
	})

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/404")
	})

	rg := r.Group("/session")
	mgr.InitSessionManager()
	rg.Use(mid.SessionMiddleware(mgr.Manager))

	rg.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "session/login.html", nil)
	})
	rg.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "zhangsan" && password == "123" {
			sdobj, ok := c.Get(mgr.SessionContextName)
			if !ok {
				c.Redirect(http.StatusFound, "/session/login")
			}
			sd := sdobj.(*mgr.SessionData)
			log.Printf("[Login]session data is: %#v", sd)
			sd.Set("isLogin", true)

			c.HTML(http.StatusOK, "session/home.html", gin.H{
				"message": username,
			})
			return
		}

		c.Redirect(http.StatusFound, "/session/login")
		return
	})

	rg.GET("/index", mid.AuthMiddleware, func(c *gin.Context) {
		c.HTML(http.StatusOK, "session/index.html", gin.H{
			"message": "ok",
		})
	})

	rg.GET("/home", mid.AuthMiddleware, func(c *gin.Context) {
		c.HTML(http.StatusOK, "session/home.html", gin.H{
			"message": "ok",
		})
	})

	rg.GET("/vip", mid.AuthMiddleware, func(c *gin.Context) {
		c.HTML(http.StatusOK, "session/vip.html", gin.H{
			"message": "ok",
		})
	})

	r.Run(":9091")
}
