
-- to create database
CREATE DATABASE library_system;

-- to connect to database
\c library_system;

-- to create user

create user bharath with password 'Password@123';

-- to grant permissions
grant all privileges on database library_system to bharath;

grant select,insert,update,delete on table books to bharath;

GRANT USAGE, SELECT ON SEQUENCE users_id_seq TO bharath;

-- to create tables

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    published_date DATE,
    isbn VARCHAR(20),
    genre VARCHAR(100),
    description TEXT
);


CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);


-- insert values to check get api:

INSERT INTO books (title, author, published_date, isbn, genre, description) 
VALUES 
    ('The Great Gatsby', 'F. Scott Fitzgerald', '1925-04-10', '978-0743273565', 'Fiction', 'The story of the fabulously wealthy Jay Gatsby and his love for the beautiful Daisy Buchanan.'),
    ('To Kill a Mockingbird', 'Harper Lee', '1960-07-11', '978-0061120084', 'Fiction', 'The story of young Jean Louise "Scout" Finch growing up in the segregated Southern United States.');


 INSERT INTO users(username,password,email)values('bharath','Bharath@023','bharath@gmail.com');

 