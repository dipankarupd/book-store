package models

// flow of the program: routes -> controller -> book.go
// user interact with route.. route send control to controller where different
// fucntions are there with business logic... controller interact with db
//operation with db reside in the db
// so all the operations shoudl have diffrent functions here as well

import (
	"github.com/dipankarupd/book-store/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// this will give the structure to how the data will be stored in the db:
type Book struct {
	gorm.Model // this model will give the structure that we talked early

	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// helps to initialize the db:
func init() {
	config.Connect()
	db = config.GetDB()

	// create books table and do necessary operations. operation performed on Book struct
	db.AutoMigrate(&Book{})
}

// functions that talk with the db:

func (b *Book) CreateBook() *Book {

	// checks if the book already exists or not. If exist, update the same one
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	// running WHERE command in mysql
	db := db.Where("id=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("id=?", ID).Delete(book)
	return book
}
