package handlers

import (
	"github.com/fentezi/httpCRUD/database"
<<<<<<< HEAD
=======
	"sync"
>>>>>>> 39e2253 (First commit)
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Isbn        uint32 `json:"isbn"`
}

type Storage interface {
	GetAllBooks() ([]Book, error)
	GetBook(id int) (Book, error)
	CreateBook(book Book) error
	UpdateBook(id int, book Book) error
	DeleteBook(id int) error
}

func NewBook() *Book {
	return &Book{}
}

<<<<<<< HEAD
func (n *Book) GetBook(id int) (Book, error) {
	b := Book{} // Создаем Book заранее
=======
var mu sync.Mutex

func (n *Book) GetBook(id int) (Book, error) {
	b := Book{}
>>>>>>> 39e2253 (First commit)
	err := database.DB.QueryRow("SELECT * FROM book WHERE id = $1", id).Scan(&b.ID, &b.Title, &b.Description, &b.Author, &b.Isbn)
	return b, err
}

func (n *Book) CreateBook(book Book) error {
<<<<<<< HEAD
=======
	mu.Lock()
	defer mu.Unlock()
>>>>>>> 39e2253 (First commit)
	stmt, err := database.DB.Prepare("INSERT INTO book(title, description, author, isbn) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Description, book.Author, book.Isbn)
	return err
}

func (n *Book) DeleteBook(id int) error {
<<<<<<< HEAD
=======
	mu.Lock()
	defer mu.Unlock()
>>>>>>> 39e2253 (First commit)
	stmt, err := database.DB.Prepare("DELETE FROM book WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func (n *Book) UpdateBook(id int, book Book) error {
<<<<<<< HEAD
=======
	mu.Lock()
	defer mu.Unlock()
>>>>>>> 39e2253 (First commit)
	stmt, err := database.DB.Prepare("UPDATE book SET title=$1, description=$2, author=$3, isbn=$4 WHERE id=$5")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Description, book.Author, book.Isbn, id)
	return err
}

func (n *Book) GetAllBooks() ([]Book, error) {
	rows, err := database.DB.Query("SELECT * FROM book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		b := Book{} // Создаем Book внутри цикла
		err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.Author, &b.Isbn)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
