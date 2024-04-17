package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ISBN          string `json:"isbn"`
	Genre         string `json:"genre"`
	Description   string `json:"description"`
}

// Function to validate book data
func validateBookData(book Book) error {
	if book.Title == "" || book.Author == "" {
		return errors.New("title and author are required")
	}
	// Additional validation rules can be added here
	return nil
}

func GetBooksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve books from the database
		rows, err := db.Query("SELECT * FROM books")
		if err != nil {
			log.Println("Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var books []Book
		for rows.Next() {
			var book Book
			if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.ISBN, &book.Genre, &book.Description); err != nil {
				log.Println("Error:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			books = append(books, book)
		}

		// Convert books slice to JSON
		json.NewEncoder(w).Encode(books)
	}
}

func GetBookHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve book ID from request parameters
		params := mux.Vars(r)
		bookID, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Println("Error:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Query the database to get the book details
		var book Book
		err = db.QueryRow("SELECT * FROM books WHERE id = $1", bookID).Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.ISBN, &book.Genre, &book.Description)
		if err != nil {
			log.Println("Error:", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Convert book to JSON
		json.NewEncoder(w).Encode(book)
	}
}

func CreateBookHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode request body into a Book struct
		var book Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			log.Println("Error:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Insert the book into the database
		_, err := db.Exec("INSERT INTO books (title, author, published_date, isbn, genre, description) VALUES ($1, $2, $3, $4, $5, $6)",
			book.Title, book.Author, book.PublishedDate, book.ISBN, book.Genre, book.Description)
		if err != nil {
			log.Println("Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateBookHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve book ID from request parameters
		params := mux.Vars(r)
		bookID, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Println("Error:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Decode request body into a Book struct
		var book Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			log.Println("Error:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Update the book in the database
		_, err = db.Exec("UPDATE books SET title=$1, author=$2, published_date=$3, isbn=$4, genre=$5, description=$6 WHERE id=$7",
			book.Title, book.Author, book.PublishedDate, book.ISBN, book.Genre, book.Description, bookID)
		if err != nil {
			log.Println("Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteBookHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve book ID from request parameters
		params := mux.Vars(r)
		bookID, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Println("Error:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Delete the book from the database
		_, err = db.Exec("DELETE FROM books WHERE id=$1", bookID)
		if err != nil {
			log.Println("Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
