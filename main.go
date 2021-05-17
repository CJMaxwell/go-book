package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book structs

type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

//Init books as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type","application/json")
	params := mux.Vars(r)

	// Loop through books and find the correct ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type","application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type","application/json")

	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {

			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	
	json.NewEncoder(w).Encode(books)
}


func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type","application/json")

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

	//Init Router
	r := mux.NewRouter()

	//Mock Book Data
	books = append(books, Book{ID: "1", Isbn: "34578", Title: "One", 
	Author: &Author{FirstName: "Maxwell", LastName: "Agbo"}})
	books = append(books, Book{ID: "2", Isbn: "40578", Title: "Two", 
	Author: &Author{FirstName: "Vincent", LastName: "Chukwudi"}})

	//Route Handlers / Endpoints

	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4500", r))
}