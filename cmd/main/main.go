// Main server
package main

import(
	// To log any error
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Rita376/MysqlConnection/pkg/routes"
)

func main(){
	// 'r' variable which is initializing the router here
	r := mux.NewRouter()
	// Register router for the function which contains all routes
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("Localhost:8000", r))
}