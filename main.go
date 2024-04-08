package main

import (
	"api/Routes"
	"log"
	"net/http"

	

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()

	routes.RegisterBookStoreRouter(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}