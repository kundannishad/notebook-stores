package main

import (
	"log"
	"net/http"
	bookRoutes "notebook-stores/pkg/routes/book"
	userRoute "notebook-stores/pkg/routes/user"

	"github.com/gorilla/mux"
)

func main() {

	//config.DbConnect()

	r := mux.NewRouter()
	bookRoutes.RegisterBookStoreRoutes(r)
	userRoute.RegisterUserRoute(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
