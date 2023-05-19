// ALL the routes are registered here

package routes

import (
	// Router package downloaded
	"github.com/gorilla/mux"
	// This is the absolute route imported for -> Getting Controller Functions
	"github.com/Rita376/MysqlConnection/pkg/controllers"
)

// Mux package here is the router used
var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}