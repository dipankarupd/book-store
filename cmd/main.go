// tells go that bookstore.route has all my routes and create my server

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dipankarupd/book-store/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	fmt.Println("Starting the server at port 9010:")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
