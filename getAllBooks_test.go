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

// func TestUpdateBookHandler(t *testing.T) {
// 	// Set up a test database or use an in-memory database
// 	// and create a new test server with your router
// 	var Bookcontroller controllers.Bookcontrollers
// 	Bookcontroller.CreateDatabase()

// 	//routing using Gollira Mux
// 	r := mux.NewRouter()
// 	r.HandleFunc("/books/{id}", Bookcontroller.UpdateBook).Methods("PUT")

// 	testServer := httptest.NewServer(r)
// 	defer testServer.Close()

// 	// Insert a sample book record into the database with the ID that you want to update
// 	// You can use a separate function to insert the sample book record into the database

// 	// Define the book payload to update an existing book record
// 	bookIDToUpdate := 1
// 	updatedBook := struct {
// 		Title           string `json:"title"`
// 		Author          string `json:"author"`
// 		PublicationYear int    `json:"publicationYear"`
// 	}{
// 		Title:           "Updated Book Title",
// 		Author:          "Updated Book Author",
// 		PublicationYear: 2023,
// 	}

// 	// Convert the book payload to JSON
// 	requestBody, err := json.Marshal(updatedBook)
// 	if err != nil {
// 		t.Fatalf("Failed to marshal JSON: %v", err)
// 	}

// 	response, err := http.NewRequest("PUT", testServer.URL+"/books/"+strconv.Itoa(bookIDToUpdate), bytes.NewBuffer(requestBody))
// 	// Make a PUT request to update the book record
// 	// response, err := http.Put(testServer.URL+"/books/"+strconv.Itoa(bookIDToUpdate), "application/json", bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		t.Fatalf("Failed to make PUT request: %v", err)
// 	}
// 	defer response.Body.Close()

// 	// Check the response status code
// 	if response.StatusCode != http.StatusOK {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.StatusCode)
// 	}

// 	// After the PUT request, make a GET request to fetch the updated book
// 	getResponse, err := http.Get(testServer.URL + "/books/" + strconv.Itoa(bookIDToUpdate))
// 	if err != nil {
// 		t.Fatalf("Failed to make GET request: %v", err)
// 	}
// 	defer getResponse.Body.Close()

// 	// Check the response status code for the GET request
// 	if getResponse.StatusCode != http.StatusOK {
// 		t.Errorf("Expected status code %d for updated book, but got %d", http.StatusOK, getResponse.StatusCode)
// 	}

// 	// Decode the response body into a Book struct
// 	var updatedBookResponse models.Book
// 	err = json.NewDecoder(getResponse.Body).Decode(&updatedBookResponse)
// 	if err != nil {
// 		t.Fatalf("Failed to decode response body: %v", err)
// 	}

// 	// Check if the updated book matches the expected values
// 	if updatedBookResponse.Title != updatedBook.Title || updatedBookResponse.Author != updatedBook.Author || updatedBookResponse.PublicationYear != updatedBook.PublicationYear {
// 		t.Errorf("Updated book does not match the expected values")
// 	}
// }
