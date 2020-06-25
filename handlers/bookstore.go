package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/beerskunk/restapi/dtos"
	"github.com/gorilla/mux"
)

// BookStore should store and retrieve books
type BookStore struct {
	books []dtos.Book
}

// Init initializes the BookStore
func (store *BookStore) Init(books []dtos.Book) {

	if books != nil && len(books) > 0 {

		for _, item := range store.books {

			store.books = append(store.books, item)
		}

	} else {
		store.books = append(store.books, dtos.Book{ID: "1", Title: "Howdy", Author: dtos.Author{FirstName: "George", LastName: "Washington"}})
		store.books = append(store.books, dtos.Book{ID: "2", Title: "Beltway Bad Boy", Author: dtos.Author{FirstName: "George", LastName: "Washington"}})
		store.books = append(store.books, dtos.Book{ID: "3", Title: "Shazam", Author: dtos.Author{FirstName: "Baby", LastName: "Jane"}})
	}
}

// Get returns a book
func (store *BookStore) Get(w http.ResponseWriter, r *http.Request) {

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_NAME"))

	params := mux.Vars(r)
	id := params["id"]
	book := dtos.Book{Author: dtos.Author{}}
	if len(id) == 0 {
		panic("404")
	}

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	qry := fmt.Sprintf("SELECT books.id as id, books.title as title, authors.firstname as firstname, authors.lastname as lastname FROM books INNER JOIN authors on books.author_id=authors.id WHERE books.id=%s", id)
	res, err := db.Query(qry)

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	fmt.Println("We got a list of books!")

	res.Next()
	res.Scan(&book.ID, &book.Title, &book.Author.FirstName, &book.Author.LastName)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// GetAll returns all the books
func (store *BookStore) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.books)
}

// Create creates a new book
func (store *BookStore) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book dtos.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	store.books = append(store.books, book)
	json.NewEncoder(w).Encode(book)
}

// Update updates an existing book
func (store *BookStore) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book dtos.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	params := mux.Vars(r)

	for index, item := range store.books {

		if item.ID == params["id"] {
			store.books = append(store.books[:index], store.books[index+1:]...)
			book.ID = item.ID
			store.books = append(store.books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(nil)
}

// Delete should delete a book
func (store *BookStore) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range store.books {

		if item.ID == params["id"] {
			store.books = append(store.books[:index], store.books[index+1:]...)
			return
		}
	}

	json.NewEncoder(w).Encode(nil)
}
