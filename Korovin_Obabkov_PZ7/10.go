package main

import (
	"fmt"
)

func main() {
	library := Library{}
	library.AddBook(Book{"Война и мир", "Лев Толстой", "В наличии"})
	library.AddBook(Book{"Преступление и наказание", "Федор Достоевский", "В наличии"})

	library.PrintBooks()

	fmt.Println("Выдаем Войну и мир, возвращаем Идиота, сжигаем 2 том мёртвых душ")
	library.TakeBook("Война и мир")
	library.ReturnBook("Идиот")
	library.BurnBook("Мёртвые души 2 том")
	fmt.Printf("\nПроверка есть ли в наличии Толстой: %v\n\n", library.SearchByAutor("Лев Толстой"))
	library.PrintBooks()
}

type Book struct {
	Title, Author, status string
}
type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}
func (l *Library) PrintBooks() {
	for _, book := range l.Books {
		fmt.Printf("Название: %v, Автор: %v, Статус: %v\n", book.Title, book.Author, book.status)
	}
}
func (l *Library) TakeBook(title string) {
	for i, book := range l.Books {
		if book.Title == title && book.status == "В наличии" {
			book.status = "Выдан"
			l.Books[i] = book
			return
		}
	}
}
func (l *Library) ReturnBook(title string) {
	for i, book := range l.Books {
		if book.Title == title && book.status == "Выдан" {
			book.status = "В наличии"
			l.Books[i] = book
			return
		}
	}
}
func (l *Library) SearchByAutor(author string) []Book {
	var result []Book
	for _, book := range l.Books {
		if book.Author == author {
			result = append(result, book)
		}
	}
	return result
}
func (l *Library) SearchByTitle(title string) []Book {
	var result []Book
	for _, book := range l.Books {
		if book.Title == title {
			result = append(result, book)
		}
	}
	return result
}
func (l *Library) BurnBook(title string) {
	for i, book := range l.Books {
		if book.Title == title {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			return
		}
	}
}
