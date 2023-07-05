package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mihir23chauhan/api/controllers"
)

func main() {

	var Bookcontroller controllers.Bookcontrollers
	Bookcontroller.CreateDatabase()

	//routing using Gollira Mux
	r := mux.NewRouter()
	r.HandleFunc("/", Bookcontroller.ServeHome).Methods("GET")
	r.HandleFunc("/books", Bookcontroller.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", Bookcontroller.GetBook).Methods("GET")
	r.HandleFunc("/books", Bookcontroller.InsertBook).Methods("POST")
	r.HandleFunc("/books/{id}", Bookcontroller.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/{id}", Bookcontroller.UpdateBook).Methods("PUT")

	log.Fatal(http.ListenAndServe(":4000", r))

}
