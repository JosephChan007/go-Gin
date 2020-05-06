package main

import (
	"github.com/gin-gonic/gin"
)

func indexAction(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "index page...",
	})
}

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!!",
		})
	})

	router.GET("/index", indexAction)

	router.Run(":9091")
}
