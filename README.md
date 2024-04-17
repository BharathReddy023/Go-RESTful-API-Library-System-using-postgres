# Go RESTful API for Simple Library System

This project implements a RESTful API for managing a library system using Go. Initially, the data was managed in memory, but it has been upgraded to use a PostgreSQL database for persistent storage. This change enhances the reliability and scalability of the system.

## Setup

### Prerequisites
- Go 1.13 or higher installed on your machine. You can download and install it from [here](https://golang.org/dl/).

### Installation
1. Clone this repository to your local machine:
git clone git@github.com:git@github.com:BharathReddy023/Go-RESTful-API-Library-System-using-postgres.git

2. Navigate to the project directory:
cd Go-RESTful-API-for-Simple-Library-System/go_restfulAPI
3. Initialize the Go module:
    go mod init git@github.com:git@github.com:BharathReddy023/Go-RESTful-API-Library-System-using-postgres.git
 
4.   Install dependencies:

   go get -u github.com/gorilla/mux
5. Run the Server:
   To run the GoLibraryAPI server, execute the following command:

   go run main.go

6. make sure postgres is installed and running 


## Creating the API

### HTTP Server Setup
1. Create a new Go file named `main.go` in your project directory.
2. Set up the HTTP server using Go's `net/http` package to listen for requests on a specified port.

## Database Integration
The API now connects to a PostgreSQL database to store and retrieve book information. This ensures that data persists across server restarts and allows for better data management and scalability.

### Request Routing
1. Implement routing using a router package like `gorilla/mux` to handle different API endpoints.
2. Define routes for CRUD operations on books:
- `GET /books`: Retrieves a list of all books.
- `GET /books/{id}`: Retrieves details of a specific book by ID.
- `POST /books`: Creates a new book.
- `PUT /books/{id}`: Updates an existing book by ID.
- `DELETE /books/{id}`: Deletes a book by ID.

## CRUD Operations
The API provides endpoints for performing CRUD operations on books, including:

Create Book: Endpoint for adding a new book to the database.
Read Books: Endpoint for retrieving a list of all books from the database.
Read Book by ID: Endpoint for retrieving details of a specific book by its ID.
Update Book: Endpoint for updating an existing book in the database.
Delete Book: Endpoint for deleting a book from the database.

### Error Handling
1. Implement proper error handling to manage common HTTP request issues.
2. Ensure the API responds with appropriate HTTP status codes and error messages.

## API Documentation

### Endpoints
- `GET /books`: Retrieves a list of all books.
- `GET /books/{id}`: Retrieves details of a specific book by ID.
- `POST /books`: Creates a new book.
- `PUT /books/{id}`: Updates an existing book by ID.
- `DELETE /books/{id}`: Deletes a book by ID.

### Request and Response Formats
- **GET /books**:
- Request: N/A
- Response: Array of book objects.
- **GET /books/{id}**:
- Request: N/A
- Response: Book object with the specified ID.
- **POST /books**:
- Request: JSON object with book title and author fields.
- Response: Newly created book object.
- **PUT /books/{id}**:
- Request: JSON object with updated book title and author fields.
- Response: Updated book object.
- **DELETE /books/{id}**:
- Request: N/A
- Response: Success message.

### Status Codes
- `200 OK`: Successful request.
- `201 Created`: Book created successfully.
- `400 Bad Request`: Invalid request data.
- `404 Not Found`: Book not found.
- `500 Internal Server Error`: Server error.

### Examples
- **GET /books**:
curl -X GET http://localhost:8080/books
- **GET /books/1**:
curl -X GET http://localhost:8080/books/1
- **POST /books**:
curl -X POST -d '{"title": "New Book", "author": "talentsprint"}' http://localhost:8080/books
- **PUT /books/1**:
curl -X PUT -d '{"title": "Updated Book", "author": "Jane Doe"}' http://localhost:8080/books/1
- **DELETE /books/1**:
curl -X DELETE http://localhost:8080/books/1


## Postman Collection

### Export Collection
1. Open Postman.
2. Click on "Collections" in the sidebar.
3. Hover over the collection you want to export and click on the ellipsis (three dots).
4. Select "Export" and choose the format "Collection v2.1".
5. Save the exported collection file to your desired location.

### Import Collection
1. Open Postman.
2. Click on "Import" in the top left corner.
3. Choose the exported collection file.
4. Click on "Import" to add the collection to your Postman workspace.

Now you can use the exported Postman collection to test your API endpoints easily.

### Learnings:
- **Go Fundamentals:** Mastered basics like syntax, data types, and control structures.
- **REST Principles:** Understood URL design, HTTP methods, and status codes.
- **CRUD Operations:** Implemented Create, Read, Update, Delete operations efficiently.
- **Error Handling:** Learned to manage errors effectively for a robust API.
- **Documentation:** Created comprehensive documentation for clear usage.
- **Testing:** Provided Postman collection for easy API testing.
