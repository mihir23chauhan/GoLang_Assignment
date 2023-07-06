package main

import (
	"bytes"
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
	//defer res.Body.Close()
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

func TestInsertABook(t *testing.T) {
	var Bookcontroller controllers.Bookcontrollers
	Bookcontroller.CreateDatabase()

	//routing using Gollira Mux
	r := mux.NewRouter()
	r.HandleFunc("/books", Bookcontroller.InsertBook).Methods("POST")

	testServer := httptest.NewServer(r)
	defer testServer.Close()

	newBook := struct {
		Title           string `json:"title"`
		Author          string `json:"author"`
		PublicationYear int    `json:"publicationYear"`
	}{
		Title:           "Test Book",
		Author:          "Test Author",
		PublicationYear: 2023,
	}
	//fmt.Println(newBook)

	// Convert the book payload to JSON
	requestBody, err := json.Marshal(newBook)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Make a POST request to create a new book
	response, err := http.Post(testServer.URL+"/books", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, response.StatusCode)
	}
}

func TestUpdateBookHandler(t *testing.T) {
	var Bookcontroller controllers.Bookcontrollers
	Bookcontroller.CreateDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/books/{id}", Bookcontroller.UpdateBook).Methods("PUT")

	testServer := httptest.NewServer(r)
	defer testServer.Close()

	bookIDToUpdate := 1
	updatedBook := struct {
		Title           string `json:"title"`
		Author          string `json:"author"`
		PublicationYear int    `json:"publicationYear"`
	}{
		Title:           "Updated Book Title",
		Author:          "Updated Book Author",
		PublicationYear: 2023,
	}

	// Convert the book payload to JSON
	requestBody, err := json.Marshal(updatedBook)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("PUT", testServer.URL+"/books/"+strconv.Itoa(bookIDToUpdate), bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to make PUT request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make PUT request: %v", err)
	}
	// Check the response status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.StatusCode)
	}
}
