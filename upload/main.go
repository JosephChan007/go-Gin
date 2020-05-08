package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("../template/**/*")

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("fileName")
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
		}
		filename := file.Filename
		log.Printf("upload file name is: %s", filename)
		filePath := fmt.Sprintf("../img/%s", filename)
		c.SaveUploadedFile(file, filePath)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.GET("/muli-upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "muli-upload.html", nil)
	})

	r.POST("/muli-upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
		}

		files := form.File["fileName"]
		for _, file := range files {
			filename := file.Filename
			log.Printf("upload file name is: %s", filename)
			filepath := fmt.Sprintf("../img/%s", filename)
			c.SaveUploadedFile(file, filepath)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.Run(":9091")
}
