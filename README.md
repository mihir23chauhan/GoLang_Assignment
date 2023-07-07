# Book Collection API

RESTful web service that allows users to manage a collection of books. It provides endpoints to perform CRUD operations (Create, Read, Update, Delete) on the book dataset.

The API can be run by executing the main function in the `main.go` file. It starts a server that listens for incoming `HTTP requests` on `port 4000`.


## **Requirements**
1. Go 1.20 
2. SQlite3 
    
    You can download and install it from the official SQLite website (https://www.sqlite.org/index.html).

### Set up the enviroment

1. Clone the Repository
2. Navigate to the Project Directory
3. Initialize Go Modules: In the project directory, run the following command to initialize Go modules and manage dependencies:

```
go mod init
```
4. Fetch Dependencies: Run the following command to fetch the required dependencies and ensure they are available:
```
go mod tidy
```

### Running the API
You can directy run API and strt the server by running main.go file 
``` 
go run main.go  
```
This command will compile and execute the `main.go` file, which will start the server and listen for incoming HTTP requests on `port 4000` (http://localhost:4000/).

Make sure that the necessary database file (`bookset.db`) is present in the same directory as the `main.go` file.

## **Endpoints**

### **GET/books**
Description: Retrieves all the books in the collection

Request Method: `GET`


### **GET/books/{id}**
Description: Retrieves a specific book by its ID.

Request Method: `GET`

Parameters:
    `{id}`: The ID of the book to retrieve.

### **POST /books**
Description: Inserts a new book into the collection.

Request Method: `POST`

Request Body: 
- JSON object representing the book to be inserted.
- ID auto generated. It can not be passed in body.
- Title and Author needs to be text and Publication Year an integer.
- Title must be in body to create a book.  

Example Request Body in json:
```
{
  "title": "Book Title",
  "author": "Book Author",
  "publicationYear": 2021
}
```
### **DELETE /books/{id}**

Description: Deletes a specific book from the collection.

Request Method: `DELETE`

Parameters:

- `{id}`: The ID of the book to delete.


### **PUT /books/{id}**
Description: Updates a specific book in the collection.

Request Method: `PUT`

Parameters:

- `{id}`: The ID of the book to update.
Request Body: 
- JSON object representing the updated book data.
- ID can not be passed or changed
- Title, author needs to be in test and publication year  an integer
- No need to pass all three items. Any one also be updated.

Example Request Body in json:
```
{
  "title": "Updated Book Title",
  "author": "Updated Book Author",
  "publicationYear": 2023
}
```
