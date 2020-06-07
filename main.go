package main

import (
	"fmt"
	"log"
	"net/http"

	perf "restapi/decorators"
	myapi "restapi/handlers"
	"restapi/interfaces"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

func main() {

	r := mux.NewRouter()
	var store interfaces.ICrud
	store = initBookStore()
	conn, ch := initRMQ()
	if conn == nil || ch == nil {
		panic("RMQ Connection failed...")
	}

	defer conn.Close()
	defer ch.Close()

	// a dumb comment
	r.HandleFunc("/api/books", perf.RestPerf(store.GetAll, ch)).Methods("GET")
	r.HandleFunc("/api/books/{id}", perf.RestPerf(store.Get, ch)).Methods("GET")
	r.HandleFunc("/api/books", perf.RestPerf(store.Create, ch)).Methods("POST")
	r.HandleFunc("/api/books/{id}", perf.RestPerf(store.Update, ch)).Methods("PUT")
	r.HandleFunc("/api/books/{id}", perf.RestPerf(store.Delete, ch)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func initBookStore() *myapi.BookStore {

	store := myapi.BookStore{}
	store.Init(nil)
	return &store
}

func initRMQ() (*amqp.Connection, *amqp.Channel) {
	queue := "TestQueue"
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	q, err := ch.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(q)

	return conn, ch
}
