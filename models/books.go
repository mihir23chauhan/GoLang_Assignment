package models

//using sqlite3 for database management

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// database file
const file string = "./bookset.db"

type Bookset struct {
	DB         *sql.DB
	TotalBooks int
}

type Book struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publicationyear"`
}

// checking if the books table is already created
// title can not be empty
const createTable string = `
CREATE TABLE IF NOT EXISTS books (
id INTEGER NOT NULL PRIMARY KEY,
title TEXT,
author TEXT,
publicationYear INTEGER
);`

// connect to the database
func ConnectDatabase() (*Bookset, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	// checking if the table already exist in the bookset.db or not
	_, err = db.Exec(createTable)

	if err != nil {
		return nil, err
	}

	// to count how many entries are already inside the table
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)
	ReportError(err)

	return &Bookset{DB: db, TotalBooks: count}, nil
}

// to generate unique id every time.
// getting the max id that exists and adding one init to get the unique id
// id will be like
// 1, 2, 3, 4, ....
func (b *Bookset) Uid() int {

	if b.TotalBooks == 0 {
		return 1
	}

	var id int
	query := "SELECT ID FROM books WHERE ID = (SELECT MAX(ID)  FROM books);"

	err := b.DB.QueryRow(query).Scan(&id)
	ReportError(err)

	return id + 1
}

// Checking if id is already present in the dataset or not.
func (b *Bookset) CheckIDinDB(id string) bool {
	var count int
	query := "SELECT COUNT(*) FROM books WHERE id = ?"
	err := b.DB.QueryRow(query, id).Scan(&count)
	ReportError(err)

	if count > 0 {
		return true
	} else {
		return false
	}
}

func ReportError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
