package routes

import (
	"github.com/gorilla/mux"
)

var HandleBookstoreRoutes = func(router *mux.Router) {
	router.NewRoute("/book", controllers.createBook).Methods("POST")
	router.NewRoute("/book", controllers.getAllBooks).Methods("GET")
	router.NewRoute("/book/{id}", controllers.getBookById).Methods("GET")
	router.NewRoute("/update/{bookId}", controllers.updateBookById).Methods("PUT")
	router.NewRoute("/delete/{bookId}", controllers.deleteBookById).Methods("DELETE")
}
