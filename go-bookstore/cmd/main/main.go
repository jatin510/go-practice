package main

import (
	"log"
	"net/http"
	"practice/go-practice/go-bookstore/pkg/routes"

	"github.com/gorilla/mux"
	"github.com/jatin510/go-bookstore/package/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("locahost:8080", r))
}
