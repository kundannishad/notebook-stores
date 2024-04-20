package book

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notebook-stores/pkg/models/book"
	utilsBook "notebook-stores/pkg/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook book.Book

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var bookResponse Response

func GetBooks(w http.ResponseWriter, r *http.Request) {

	newBooks := book.GetAllBooks()
	bookResponse = Response{Message: "Book Retrieved Successfully", Data: newBooks}
	bookResponseJSON, _ := json.Marshal(bookResponse)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bookResponseJSON)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing !")
	}

	bookDetails, _ := book.GetBookById(Id)

	bookResponse := Response{Message: "Book Details retrive sucessfully !", Data: bookDetails}

	res, _ := json.Marshal(bookResponse)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	createBook := &book.Book{}
	utilsBook.ParseBody(r, createBook)

	bookRes := createBook.CreateBook()

	bookResponse = Response{Message: "Book Created successfully", Data: bookRes}
	res, _ := json.Marshal(bookRes)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error While parsingb the data")
	}
	book := book.DeleteBookById(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updateBook = &book.Book{}
	utilsBook.ParseBody(r, updateBook)

	vars := mux.Vars(r)

	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing !")
	}

	bookDetails, db := book.GetBookById(Id)

	if updateBook.Title != "" {
		bookDetails.Title = updateBook.Title
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.ISBN != "" {
		bookDetails.ISBN = updateBook.ISBN
	}

	if updateBook.Language != "" {
		bookDetails.Language = updateBook.Language
	}

	db.Save(&bookDetails)

	bookResponse = Response{Message: "Book Update Successfully !", Data: bookDetails}

	res, _ := json.Marshal(bookResponse)

	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}
