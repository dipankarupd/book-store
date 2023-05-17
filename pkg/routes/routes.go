// routing codes
package routes

import (
	"github.com/dipankarupd/book-store/pkg/controllers"
	"github.com/gorilla/mux"
)

// creating the routes:

// in go, you can assign a var to a func if the signature of var and func is same
var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{bookID}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookID}", controllers.DeleteBook).Methods("DELETE")

}
