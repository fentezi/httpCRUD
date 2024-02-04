package main

import (
	"github.com/fentezi/httpCRUD/database"
	"github.com/fentezi/httpCRUD/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	godotenv.Load()
	database.ConnectDB()
}

func main() {
	defer database.DB.Close()
	memoryStorage := handlers.NewBook()
	handle := handlers.NewHandler(memoryStorage)
	r := mux.NewRouter()
<<<<<<< HEAD
=======
	log.Printf("Запуск сервера!")
>>>>>>> 39e2253 (First commit)
	r.HandleFunc("/books", handle.GetAllBooks).Methods("GET")
	r.HandleFunc("/book/{id}", handle.GetBook).Methods("GET")
	r.HandleFunc("/book", handle.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", handle.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", handle.DeleteBook).Methods("DELETE")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
