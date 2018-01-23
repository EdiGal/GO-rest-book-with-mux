package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct (Modal)
type Book struct {
	ID     string  "json:\"id\""
	Isbn   string  "json:\"isbn\""
	Title  string  "json:\"title\""
	Author *Author "json:\"author\""
}

//Author struct
type Author struct {
	Firstname string "json:\"firstname\""
	Lastname  string "json:\"lastname\""
}

//Mock data
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	// json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()
	// rand.Int()
	// strconv.ParseInt(3)
	//Mock data
	books = append(books, Book{ID: "1", Isbn: "999999", Title: "Fire and Fury", Author: &Author{Firstname: "Michael", Lastname: "Wolff"}})
	books = append(books, Book{ID: "2", Isbn: "888888", Title: "Milk and Honey", Author: &Author{Firstname: "Rupi", Lastname: "Kaur"}})
	books = append(books, Book{ID: "3", Isbn: "777777", Title: "A Beautiful Mind, a Beautiful Life", Author: &Author{Firstname: "Michael", Lastname: "Wolff"}})
	books = append(books, Book{ID: "4", Isbn: "666666", Title: "The Sun and Her Flowers", Author: &Author{Firstname: "Rupi", Lastname: "Kaur"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
