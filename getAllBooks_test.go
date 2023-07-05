package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mihir23chauhan/api/controllers"
	"github.com/mihir23chauhan/api/models"
)

func TestGetAllBooks(t *testing.T) {
	var Bookcontroller controllers.Bookcontrollers
	Bookcontroller.CreateDatabase()

	//routing using Gollira Mux
	r := mux.NewRouter()
	r.HandleFunc("/books", Bookcontroller.GetAllBooks).Methods("GET")

	testServer := httptest.NewServer(r)
	defer testServer.Close()

	res, err := http.Get(testServer.URL + "/books")

	if err != nil {
		t.Fatalf("Failed with all books get request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.StatusCode)
	}
	var books []models.Book
	err = json.NewDecoder(res.Body).Decode(&books)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}
	fmt.Println(books)

}

func TestGetABookWithPresentID(t *testing.T) {
	var Bookcontroller controllers.Bookcontrollers
	Bookcontroller.CreateDatabase()

	//routing using Gollira Mux
	r := mux.NewRouter()
	r.HandleFunc("/books/{id}", Bookcontroller.GetBook).Methods("GET")

	testServer := httptest.NewServer(r)
	defer testServer.Close()

	res, err := http.Get(testServer.URL + "/books/1")

	if err != nil {
		t.Fatalf("Failed with all books get request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.StatusCode)
	}
	var books models.Book
	err = json.NewDecoder(res.Body).Decode(&books)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}
	fmt.Println(books)

}

func TestGetABookNoID(t *testing.T) {
	var Bookcontroller controllers.Bookcontrollers
	Bookcontroller.CreateDatabase()

	//routing using Gollira Mux
	r := mux.NewRouter()
	r.HandleFunc("/books/{id}", Bookcontroller.GetBook).Methods("GET")

	testServer := httptest.NewServer(r)
	defer testServer.Close()

	res, err := http.Get(testServer.URL + "/books/100")

	if err != nil {
		t.Fatalf("Failed with all books get request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.StatusCode)
	}
	var IdNotFound string
	err = json.NewDecoder(res.Body).Decode(&IdNotFound)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}
	fmt.Println(IdNotFound)

}

func TestDeleteABook(t *testing.T) {
	var Bookcontroller controllers.Bookcontrollers
	Bookcontroller.CreateDatabase()

	//routing using Gollira Mux
	r := mux.NewRouter()
	r.HandleFunc("/books/{id}", Bookcontroller.DeleteBook).Methods("DELETE")

	testServer := httptest.NewServer(r)
	defer testServer.Close()
	bookIDToDelete := 4

	req, err := http.NewRequest("DELETE", testServer.URL+"/books/"+strconv.Itoa(bookIDToDelete), nil)

	if err != nil {
		t.Fatalf("Failed with delete the record: %v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		t.Fatalf("Failed with delete the record: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.StatusCode)
	}
	var deleted string
	err = json.NewDecoder(res.Body).Decode(&deleted)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}
	fmt.Println(deleted)

}
