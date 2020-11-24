package main

type Book struct {
	Id int`db:"id"`
	BookName string`db:"book_name"`
	Price float64`db:"price"`
}
