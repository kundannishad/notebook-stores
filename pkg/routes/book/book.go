package book

import (
	"notebook-stores/pkg/controllers/book"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", book.GetBooks).Methods("GET")
	router.HandleFunc("/book", book.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", book.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", book.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookId}", book.DeleteBookById).Methods("DELETE")

}
