package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// Book is a placeholder for a book
type Book struct {
	Id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite", "./books.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	// Create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("Error creating table")
		log.Println(err)
	} else {
		log.Println("Successfully created table books")
	}
	statement.Exec()
	dbOperations(db)
}

func dbOperations(db *sql.DB) {
	// Create
	statement, _ := db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
	statement.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	log.Println("Inserted the book into database!")

	// Read
	rows, _ := db.Query("SELECT id, name, author FROM books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.Id, &tempBook.name, &tempBook.author)
		log.Printf("ID: %d, Book: %s, Author: %s\n", tempBook.Id, tempBook.name, tempBook.author)
	}

	// Update
	statement, _ = db.Prepare("update books set name=? where id=?")
	statement.Exec("The Tale of Two Cities", 1)
	log.Println("Successfully updated the book in database!")

	// Delete
	statement, _ = db.Prepare("delete from books where id=?")
	statement.Exec(1)
	log.Println("Successfully deleted the book in database!")
}
