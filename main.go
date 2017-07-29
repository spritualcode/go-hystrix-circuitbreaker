package main

import (
	"log"
	"net/http"

	"github.com/spritualcode/go-hystrix-circuitbreaker/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := NewRouter()
	log.Println("Starting up...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	AddRoutes(router)
	return router
}

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/example", handler.TestingHandler)
}
