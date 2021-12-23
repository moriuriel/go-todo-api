package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}

func main() {

	router := createRouter()
	log.Fatal(http.ListenAndServe(":5000", router))

	fmt.Println("Server is starting on port")
}
