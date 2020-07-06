package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	perf "github.com/beerskunk/restapi/src/decorators"
	myapi "github.com/beerskunk/restapi/src/handlers"
	"github.com/beerskunk/restapi/src/interfaces"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/streadway/amqp"
)

func main() {

	r := mux.NewRouter()
	var store interfaces.IBookStore
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

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*", "http://localhost:3000"},
		AllowedHeaders: []string{
			"Accept",
			"Accept-Encoding",
			"Accept-Language",
			"Cache-Control",
			"Connection",
			"DNT",
			"Host",
			"Origin",
			"Pragma",
			"Referer",
			"User-Agent",
		},
		AllowedMethods: []string{
			"DELETE",
			"GET",
			"OPTIONS",
			"POST",
			"PUT",
		},
	})

	log.Fatal(http.ListenAndServe(":8080", corsOpts.Handler(r)))
}

func initBookStore() *myapi.BookStore {

	store := myapi.BookStore{}
	store.Init()
	return &store
}

func initRMQ() (*amqp.Connection, *amqp.Channel) {
	queue := os.Getenv("RMQ_QNAME")
	uname := os.Getenv("RMQ_UNAME")
	pwd := os.Getenv("RMQ_PWD")
	domain := os.Getenv("RMQ_URL")

	if queue == "" {
		panic("RMQ Queue Name is empty")
	}

	if uname == "" {
		panic("RMQ Username is empty")
	}

	if pwd == "" {
		panic("RMQ Pwd is empty")
	}

	if domain == "" {
		panic("RMQ Domain is empty")
	}

	rmq := fmt.Sprintf("amqp://%s:%s@%s/", uname, pwd, domain)
	conn, err := amqp.Dial(rmq)
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
