package model

type Book struct {
	Id    int64   `db:"id" form:"id"`
	Name  string  `db:"name" form:"name"`
	Price float64 `db:"price" form:"price"`
}
