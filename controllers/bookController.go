package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihir23chauhan/api/models"
)

type Bookcontrollers struct{}

var bookset *models.Bookset

func (c *Bookcontrollers) CreateDatabase() {
	bookset, _ = models.ConnectDatabase()
}

// GET/
func (c *Bookcontrollers) ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Welcome to Book Collection</h1>
	<p>GET/books			to get all the collection of books</p>
	<p>GET/books/{id}		to get a collection of book with given id</p>
	<p>POST/books			to insert the book in database</p>
	<p>PUT/books/{id}			to update the book with given id in database</p>
	<p>DELETE/books/{id}	to delete the book with given id in database </p>

	`))

}

// GET/books
// Handler/ Controller
func (c *Bookcontrollers) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all the course")
	w.Header().Set("Content-Type", "application/json")

	if IsEmptyDataset() {
		json.NewEncoder(w).Encode("Dataset is empty")
		return
	}

	// rows = all the rows of dataset
	rows, err := bookset.DB.Query("SELECT * FROM books;")
	ReportError(err)

	// collection all the rows in slice to convert it into json
	var collection []models.Book

	for rows.Next() {
		var b models.Book
		err = rows.Scan(&b.Id, &b.Title, &b.Author, &b.PublicationYear)
		ReportError(err)

		collection = append(collection, b)
	}
	// if any error while goinig through all the rows
	err = rows.Err()
	ReportError(err)

	json.NewEncoder(w).Encode(collection)
}

// GET/books/{id}
func (c *Bookcontrollers) GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request for a book id ")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	if IsEmptyDataset() {
		json.NewEncoder(w).Encode("Dataset is empty")
		return
	}
	// check if id is present in the database
	if bookset.CheckIDinDB(id) {

		getQuery := "SELECT * FROM books WHERE id = ?"

		row := bookset.DB.QueryRow(getQuery, id)

		var book models.Book

		// extract the information out of the feteched row
		err := row.Scan(&book.Id, &book.Title, &book.Author, &book.PublicationYear)
		ReportError(err)

		// book with the given ID in json format
		json.NewEncoder(w).Encode(book)

	} else {
		fmt.Println("Id not found")
		json.NewEncoder(w).Encode("No course with the given id")
		return
	}
}

// POST/books
// expect request with a proper form of json in body
func (c *Bookcontrollers) InsertBook(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request received to insert an element")
	w.Header().Set("Content-Type", "application/json")

	book := ConvertRequestToBook(r)

	// Not valid body
	if book.Id == -1 {
		fmt.Println("No data passed inside the request")
		json.NewEncoder(w).Encode("No or not valid data passed inside the request")
		return
	} else if book.Id == -4 {
		fmt.Println("No valid type of data passed inside the request")
		json.NewEncoder(w).Encode("No valid type of data passed inside the request")
		return
	} else if book.Id == -2 {
		fmt.Println("Id was provided")
		json.NewEncoder(w).Encode("Id was provided. User is not suppposed to provide an id")
		return
	} else if book.Id == -3 {
		fmt.Println("Title was not provided.")
		json.NewEncoder(w).Encode("Title was not provided. Title is needed to insert a book")
		return
	}
	// adding to the total number of books
	bookset.TotalBooks++

	insertQuery := "INSERT INTO books VALUES (?, ?, ?,?)"
	_, err := bookset.DB.Exec(insertQuery, &book.Id, &book.Title, &book.Author, &book.PublicationYear)

	ReportError(err)

	fmt.Println("Succesfully Inserted.")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(book)

}

// DELETE/{id} | delete the record with the given id
func (c *Bookcontrollers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to delete a book id ")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	// Check is id exists or not in dataset
	if bookset.CheckIDinDB(id) {
		getQuery := "DELETE FROM books WHERE id = ?"

		_, err := bookset.DB.Exec(getQuery, id)

		if err != nil {
			json.NewEncoder(w).Encode("Error")
			log.Fatal(err)
			return
		}
		bookset.TotalBooks--
		json.NewEncoder(w).Encode("Record deleted.")
	} else {
		json.NewEncoder(w).Encode("No record with the given id.")
		fmt.Println("No record with the given id.")
	}
}

// PUT/books/{id} | update the existed record
func (c *Bookcontrollers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request for an update")
	w.Header().Set("Content-Type", "application/json")

	book := ConvertRequestToBook(r)
	//fmt.Println(book)

	if book.Id == -1 {
		fmt.Println("No data passed inside the request")
		json.NewEncoder(w).Encode("No or not valid data passed inside the request")
		return
	} else if book.Id == -4 {
		fmt.Println("No valid type of data passed inside the request")
		json.NewEncoder(w).Encode("No valid type of data passed inside the request")
		return
	} else if book.Id == -2 {
		fmt.Println("Id was provided")
		json.NewEncoder(w).Encode("Id was provided. User is not suppposed to provide an id")
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	if bookset.CheckIDinDB(id) {

		if book.Title != "" && book.Author != "" && book.PublicationYear != 0 {
			getQuery := "UPDATE books SET title = ?, author = ?, publicationyear =? WHERE id = ?"

			_, err := bookset.DB.Exec(getQuery, book.Title, book.Author, book.PublicationYear, id)
			ReportError(err)
			json.NewEncoder(w).Encode("Record updated.")

			return
		}

		if book.Title != "" {
			getQuery := "UPDATE books SET title = ? WHERE id = ?"

			_, err := bookset.DB.Exec(getQuery, book.Title, id)
			ReportError(err)
		}
		if book.Author != "" {
			getQuery := "UPDATE books SET author = ? WHERE id = ?"

			_, err := bookset.DB.Exec(getQuery, book.Author, id)
			ReportError(err)
		}
		if book.PublicationYear != 0 {
			getQuery := "UPDATE books SET publicationyear =? WHERE id = ?"

			_, err := bookset.DB.Exec(getQuery, book.PublicationYear, id)
			ReportError(err)
		}

		json.NewEncoder(w).Encode("Record updated.")
	} else {
		json.NewEncoder(w).Encode("No record with the given id.")
		fmt.Println("No record with the given id.")
	}
}

// Check where the request data is in correct form or not.
// Return Book models.Book
// book.Id > 0, book good to add in dataset
// book.Id == -1, Not data inside body
// bood.ID == -2, Id is provided
// book.ID == -3, title was not provided
// book.ID == -4 , passed datatype was not correct
func ConvertRequestToBook(r *http.Request) models.Book {

	var book models.Book
	var jsonData map[string]interface{}

	// Parse JSON data from request body into the map
	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		if err == io.EOF {
			fmt.Println("No body was passed")
			book.Id = -1
			return book
		}
	}
	_, isId := jsonData["id"]
	title, isTitle := jsonData["title"]
	author, isAuthor := jsonData["author"]
	publicationYear, isPublicationYear := jsonData["publicationyear"]

	if !isTitle && !isAuthor && !isPublicationYear {
		fmt.Println("No Data was passed")
		book.Id = -1
		return book
	}

	if isId {
		book.Id = -2
		return book
	}
	if !isTitle {
		fmt.Println("Title is not passed")
		book.Id = -3
	} else {
		if t, ok := title.(string); ok {
			book.Id = bookset.Uid()
			book.Title = t
		} else {
			book.Id = -4
			fmt.Println("Not valid type")
			return book
		}
	}
	if isAuthor {
		if a, ok := author.(string); ok {
			book.Author = a
		} else {
			book.Id = -4
			fmt.Println("Not valid type")
			return book
		}
	}

	if isPublicationYear {
		fmt.Printf("%T", publicationYear)
		if y, ok := publicationYear.(float64); ok {
			book.PublicationYear = int(y)
		} else {
			book.Id = -4
			fmt.Println("Not valid type")
			return book
		}
	}

	return book
}

func IsEmptyDataset() bool {
	if bookset.TotalBooks == 0 {
		fmt.Println("Dataset is empty")
		return true
	}
	return false
}

func ReportError(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
