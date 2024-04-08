package routes

import (
	"api/controllers"

	"github.com/gorilla/mux"
)

// RegisterBookStoreRouter registers bookstore route
func RegisterBookStoreRouter(router *mux.Router) {

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")

	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")

	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")

	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}


