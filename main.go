package main

import (
    "database/sql"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
    "apis/handlers"
)

func main() {
    // PostgreSQL connection string
    connStr := "user=bharath dbname=library_system password=Password@123 host=localhost sslmode=disable"

    // Connect to the PostgreSQL database
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create a new router
    router := mux.NewRouter()

    // Define routes using handler functions from the handlers package
    router.HandleFunc("/books", handlers.GetBooksHandler(db)).Methods("GET")
    router.HandleFunc("/books/{id}", handlers.GetBookHandler(db)).Methods("GET")
    router.HandleFunc("/books", handlers.CreateBookHandler(db)).Methods("POST")
    router.HandleFunc("/books/{id}", handlers.UpdateBookHandler(db)).Methods("PUT")
    router.HandleFunc("/books/{id}", handlers.DeleteBookHandler(db)).Methods("DELETE")

    // Start the HTTP server
    log.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
