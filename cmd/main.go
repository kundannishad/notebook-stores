package main

import (
	"log"
	"net/http"
	bookRoutes "notebook-stores/pkg/routes/book"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	bookRoutes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
