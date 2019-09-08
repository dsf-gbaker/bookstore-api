package main

import (
	"log"
	"net/http"

	myapi "restapi/handlers"
	"restapi/interfaces"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	var store interfaces.ICrud
	store = initBookStore()

	r.HandleFunc("/api/books", store.GetAll).Methods("GET")
	r.HandleFunc("/api/books/{id}", store.Get).Methods("GET")
	r.HandleFunc("/api/books", store.Create).Methods("POST")
	r.HandleFunc("/api/books/{id}", store.Update).Methods("PUT")
	r.HandleFunc("/api/books/{id}", store.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func initBookStore() *myapi.BookStore {

	store := myapi.BookStore{}
	store.Init(nil)
	return &store
}
