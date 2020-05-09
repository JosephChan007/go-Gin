package dao

import (
	"fmt"
	. "github.com/JosephChan007/go-Gin/BMS/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func InitDB() (err error) {
	dsn := "root:123456@tcp(hdfs-host4:3306)/test"

	// 注意：此处db赋值不能用:=，否则db有效作用域仅能在InitDB()方法内，其他方法都为nil
	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func GetBook(id int64) (book *Book, err error) {
	book = &Book{}
	sql := "select id, name, price from book where id=?"
	err = db.Get(book, sql, id)
	if err != nil {
		log.Printf("DB query a book failed, err:%v\n", err)
		return
	}
	return
}

func GetBookList() (list []*Book, err error) {
	sql := "select id, name, price from book"
	err = db.Select(&list, sql)
	if err != nil {
		log.Printf("DB query book list failed, err:%v\n", err)
		return
	}
	return
}

func AddBook(book *Book) (err error) {
	log.Printf("book param is: %v\n", book)
	sql := "insert into book (name, price) values (?, ?)"
	res, err := db.Exec(sql, book.Name, book.Price)
	if err != nil {
		log.Printf("DB add a book failed, err:%v\n", err)
		return
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Printf("DB add a book failed, err:%v\n", err)
		return
	}
	log.Printf("DB add a book success, book id is:%d\n", lastId)
	return
}

func DeleteBook(id int64) (err error) {
	sql := "delete from book where id = ?"
	_, err = db.Exec(sql, id)
	if err != nil {
		log.Printf("DB delete a book failed, err:%v\n", err)
		return
	}
	return
}

func UpdateBook(book *Book) (err error) {
	sql := "update book set name=?, price=? where id = ?"
	_, err = db.Exec(sql, book.Name, book.Price, book.Id)
	if err != nil {
		log.Printf("DB update a book failed, err:%v\n", err)
		return
	}
	return
}
