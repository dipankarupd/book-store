//after routing all the controls come here

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dipankarupd/book-store/pkg/models"
	"github.com/dipankarupd/book-store/pkg/utils"
	"github.com/gorilla/mux"
)

// getting the book struct from book.go
var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()

	// convert everything to json and stored in res
	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) // statusOK gives us 200
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {

	// get the required id from the request
	vars := mux.Vars(r)
	bookId := vars["bookID"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error occured while requesting for the ID")
	}

	// models.GetBookByID returns a book and db var, I dont want to use the
	// db variable, so I used _
	bookDetails, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	// ready for unmarshalling the req. Send to util.go which does that
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)

	b := CreateBook.CreateBook()

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId := vars["bookID"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error occured..")
	}
	book := models.DeleteBook(ID)

	// sending the response:
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookID := vars["bookID"]

	ID, err := strconv.ParseInt(bookID, 0, 0)

	if err != nil {
		fmt.Println("Error while Parsing")
	}

	// get the book by id:
	bookDetail, db := models.GetBookByID(ID)

	// check if the book exists in our db or not
	if updateBook.Name != "" {
		bookDetail.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetail.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetail.Publication = updateBook.Publication
	}

	db.Save(&bookDetail)

	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
