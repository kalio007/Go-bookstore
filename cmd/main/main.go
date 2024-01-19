package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kalio007/Go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	//ref: for gin it was gin.default
	routes.RegisterBookStoreRoutes(r)
	//mostly using http to manage the inital route and server
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
