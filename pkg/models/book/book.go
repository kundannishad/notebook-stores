package book

import (
	"notebook-stores/pkg/config"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	ID              int       `gorm:"primaryKey" json:"id,omitempty"`
	Title           string    `gorm:"not null" json:"title,omitempty"`
	Author          string    `gorm:"not null" json:"author,omitempty"`
	PublicationYear int       `json:"publication_year,omitempty"`
	ISBN            string    `json:"isbn,omitempty"`
	Genre           string    `json:"genre,omitempty"`
	Language        string    `json:"language,omitempty"`
	Publisher       string    `json:"publisher,omitempty"`
	Pages           int       `json:"pages,omitempty"`
	Price           float64   `json:"price,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

func init() {
	config.DbConnect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBookById(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
