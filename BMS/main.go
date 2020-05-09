package main

import (
	"github.com/JosephChan007/go-Gin/BMS/dao"
	. "github.com/JosephChan007/go-Gin/BMS/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		panic(err)
	}
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.LoadHTMLGlob("../template/**/*")
	r.Static("/static", "../img") // 参数1：URL代码中的文件前缀；参数：实际存放文件路径

	rg := r.Group("/book")

	rg.GET("/list", func(c *gin.Context) {
		list, _ := dao.GetBookList()
		c.HTML(http.StatusOK, "book/list.html", gin.H{
			"data": list,
		})
	})

	rg.GET("/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "book/book.html", nil)
	})

	rg.POST("/new", func(c *gin.Context) {
		nameVal := c.PostForm("name")
		price := c.PostForm("price")
		priceVal, _ := strconv.ParseFloat(price, 64)
		log.Printf("book param is name(%s), price(%f)\n", nameVal, priceVal)
		book := &Book{
			Name:  nameVal,
			Price: priceVal,
		}
		_ = dao.AddBook(book)
		c.Redirect(http.StatusMovedPermanently, "/book/list")
	})

	rg.GET("/update", func(c *gin.Context) {
		idstr := c.Query("id")
		idVal, _ := strconv.ParseInt(idstr, 10, 64)
		log.Printf("book request param is id(%d)\n", idVal)
		book, _ := dao.GetBook(idVal)
		log.Printf("book param is: %v\n", book)
		c.HTML(http.StatusOK, "book/book.html", gin.H{
			"id":    book.Id,
			"name":  book.Name,
			"price": book.Price,
		})
	})

	rg.POST("/update", func(c *gin.Context) {
		id := c.PostForm("id")
		idVal, _ := strconv.ParseInt(id, 10, 64)
		nameVal := c.PostForm("name")
		price := c.PostForm("price")
		priceVal, _ := strconv.ParseFloat(price, 64)
		log.Printf("book param is id(%d), name(%s), price(%f)\n", idVal, nameVal, priceVal)
		book := &Book{
			Id:    idVal,
			Name:  nameVal,
			Price: priceVal,
		}
		_ = dao.UpdateBook(book)
		c.Redirect(http.StatusMovedPermanently, "/book/list")
	})

	rg.GET("/delete", func(c *gin.Context) {
		id := c.Query("id")
		idVal, _ := strconv.ParseInt(id, 10, 64)
		log.Printf("book param is id(%d)\n", idVal)
		_ = dao.DeleteBook(idVal)
		c.Redirect(http.StatusMovedPermanently, "/book/list")
	})

	r.Run(":9091")

}
