# Book Collection API

RESTful web service that allows users to manage a collection of books. It provides endpoints to perform CRUD operations (Create, Read, Update, Delete) on the book dataset.



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
- ID auto generated. It should not be passed in body.
- Title and Author needs to be text and Publication Year an integer. 

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
