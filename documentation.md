# Documenation of controllers and model
You can read about how to set up the enviroment and run the api in README.md file

## **Models**

### **Bookset** 
The Bookset struct represents the book collection and provides methods to interact with the database.

```
type Bookset struct {
	DB         *sql.DB
	TotalBooks int
}
```

#### `ConnectDatabase() (*Bookset, error)`
- Description: Establishes a connection to the book database and initializes the Bookset object.

- This function connects to the SQLite database file (`bookset.db`) using the `sql.Open` function and checks if the `books table` already exists in the database. 
- If the table does not exist, it creates it using the `CREATE TABLE IF NOT EXISTS` SQL statement.

#### `Uid() int`
- Description: Generates a unique ID for a book.

- Generates a unique ID for a book by retrieving the `maximum ID `from the books table and `adding one` to it. 
- If no books exist in the table, it `returns 1 as the initial ID`. 

#### `CheckIDinDB(id string) bool`
- Description: Checks if a book ID exists in the database, by executing a `SELECT COUNT(*) FROM books WHERE id = ?` SQL statement. 
- It returns `true` if the ID exists and `false` otherwise.

### **Book**
The Book struct is used to represent individual books when retrieving, inserting, updating, or deleting them from the collection.
```
type Book struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publicationyear"`
}
```
## **Controllers**

The `Bookcontrollers` struct provides methods to handle HTTP requests and perform CRUD operations on the book collection. It uses the Gorilla Mux package for routing and the `encoding/json` package for JSON encoding and decoding.

### `CreateDatabas()`
- Description: Creates a connection to the book database and initializes the Bookset object as `bookset *models.Bookset`. 

- Uses `models.ConnectDatabase() ` method in `models/Book.go` file. 

### `ServeHome(w http.ResponseWriter, r *http.Request)`
- Description: Handles the home page request and displays information about the available endpoints.

### `GetAllBooks(w http.ResponseWriter, r *http.Request)`

- Description: Retrieves all the books in the collection and returns them as a `JSON array`.

- This method handles the `GET` request to `"/books"` endpoint. 

- If the book dataset is empty (checked using the `IsEmptyDataset` function), it encodes and returns a message indicating that the dataset is empty.
- Otherwise, it retrieves all the rows from the books table in the database using the `SELECT SQL statement` and populates a `slice of models.Book` structs with the fetched data. The slice is then encoded as JSON and sent as the response.


### `GetBook(w http.ResponseWriter, r *http.Request)`
- Description: Retrieves a specific book by its `ID` and returns it as a `JSON object`.

- This method handles the `GET` request to `"/books/{id}"` endpoint. It extracts the book ID from the URL parameters using `mux.Vars(r)`. 

- If the book dataset is empty (checked using the `IsEmptyDataset` function), it encodes and returns a message indicating that the dataset is empty. 

- It then checks if the book ID exists in the database using the `CheckIDinDB` method of the `Bookset struct`. 

- If the ID is present, it executes a `SELECT WHERE ID = ? SQL statement` to fetch the book data from the books table in dataset.

- Finally, the book struct is encoded as `JSON` and sent as the response.

### `ConvertRequestToBook(r *http.Request) models.Book`
- Description: This function takes `*HTTP request` as input and converts the `JSON data` in the request body to a `models.Book struct`. 

- It checks for the presence of required fields (title), validates the data types, and returns the book struct. 

    - Title must be present while inserting a book. 
    - ID can not be present in any request. 
    - Title, Author must be Text and Publicatoin Year must be an Integer.

- If any errors occur during the conversion, `appropriate error codes (-1 to -4)` are set in the book struct to indicate the type of error. Check `retuned Book.ID` for error type.
```
If book.ID == -1, 
    no valid data was passed in the request body
else if book.ID == -2
    ID was provided in the request
else if book.ID == -3   
    the title was not provided
else if book.ID == -4 
    data types in the request body are not valid
```
### `InsertBook(ResponseWriter, r *http.Request)`

- Description: Inserts a new book into the collection based on the data provided in the request body.

- This method handles the `POST request` to "`/books`" endpoint. 
- It first checks if the request body contains valid JSON data and extracts the book information using the `ConvertRequestToBook` function. 
- If the book data is valid, it increments the `TotalBooks count` in the `Bookset struct` and inserts the book into the books table in the database using the `INSERT INTO books VALUES(,,,,)` SQL statement. 
- Finally, the inserted book is encoded as `JSON` and returned in the response body with a `status code of 201 (Created)`.


### `UpdateBook(w http.ResponseWriter, r *http.Request)`
- Description: Updates a specific book in the collection based on its ID and the updated data provided in the request body.

- This method handles the `PUT` request to `"/books/{id}"` endpoint. 
- It extracts the book ID from the URL parameters using `mux.Vars(r)`. 
- It checks if the book ID exists in the database using the `CheckIDinDB` method of the Bookset struct. 
- If the ID is present, it extracts the updated book data from the request body using the `ConvertRequestToBook` function.

- The method then checks if any specific fields (`title, author, publication year`) were updated in the request body. 
- If any field is provided, it executes an appropriate `UPDATE book SET  Where id = ? ` SQL statement to update the corresponding field in the books table.
- Finally, a success message is returned in the response body.

### `DeleteBook(w http.ResponseWriter, r *http.Request)`
- Description: Deletes a specific book from the collection based on its ID.

- This method handles the `DELETE` request to `"/books/{id}"` endpoint. 
- It extracts the book ID from the URL parameters using `mux.Vars(r)`. 
- It checks if the book ID exists in the database using the `CheckIDinDB` method of the Bookset struct. 
- If the ID is present, it executes a `DELETE FROM books WHERE id = ?` SQL statement to remove the book from the books table. 
- The TotalBooks count in the Bookset struct is decremented, and a success message is returned in the response body. 
- If the book ID is not found in the database, an appropriate error message is returned.

### `IsEmptyDataset() bool`
- Description: Checks if the book dataset is empty.

### `ReportError(err error)`
- Description: Reports any errors that occur during database operations.