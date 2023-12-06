package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/beerskunk/restapi/src/dtos"
	"github.com/gorilla/mux"
)

// BookStore should store and retrieve books
type BookStore struct {
	db *sql.DB
}

// Init initializes the BookStore
func (store *BookStore) Init() {

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}

	store.db = db
}

// Get returns a book
func (store *BookStore) Get(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	book := dtos.Book{}
	if len(id) == 0 {
		panic("404")
	}

	qry := fmt.Sprintf("SELECT books.id as id, books.title as title, authors.firstname as firstname, authors.lastname as lastname FROM books INNER JOIN authors on books.author_id=authors.id WHERE books.id=%s", id)
	res, err := store.db.Query(qry)

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	res.Next()
	res.Scan(&book.ID, &book.Title, &book.Author.FirstName, &book.Author.LastName)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// GetAll returns all the books
func (store *BookStore) GetAll(w http.ResponseWriter, r *http.Request) {

	qry := "SELECT books.id, books.title, authors.firstname, authors.lastname FROM books INNER JOIN authors on books.author_id=authors.id"
	res, err := store.db.Query(qry)
	var books []dtos.Book
	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	for res.Next() {

		var book dtos.Book
		err := res.Scan(&book.ID, &book.Title, &book.Author.FirstName, &book.Author.LastName)

		if err != nil {
			panic(err.Error())
		}

		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Create creates a new book
func (store *BookStore) Create(w http.ResponseWriter, r *http.Request) {
	var book dtos.Book
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Update updates an existing book
func (store *BookStore) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}

// Delete should delete a book
func (store *BookStore) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}
