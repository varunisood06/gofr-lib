# Gofr Project

This is a simple Go project that uses Gofr framework and provides a RESTful API for managing CRUD operations over a library management.

## Getting Started

These instructions will get you a clone of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have Go installed on your machine. You can download it from the [official website](https://golang.org/dl/).
You need to have Docker Desktop installed on your machine.

### Installing

To install the project, follow these steps:

1. Clone the repository: git clone https://github.com/varunisood06/gofr-lib.git
2. Navigate to the project directory: cd go-project
3. Download the dependencies: go get .
4. Run the following commands in a terminal to create a table in MySQL docker image
   - `docker run --name zop-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=varu -p 2021:3306 -d mysql:8.0.30`
   - `docker exec -it zop-mysql mysql -uroot -proot123 varu -e "CREATE TABLE Books (BookID INT AUTO_INCREMENT PRIMARY KEY, Title VARCHAR(255) NOT NULL, AuthorID INT NOT NULL, Issued VARCHAR(255) NOT NULL);"



## Running the tests

To run the tests, use the following command: `go test .`


## Running the application

To start the application, use the following command: `go run .`


The application will start and listen on port 9000.

## API Endpoints
- `GET/library/{id}`: Get a details of a book by id.
- `POST/book`: Add new book into DB.
- `PUT/book/{id}`: Update the details of a book.
- `DELETE/book/{id}`: Delete the details for a book.

## Built With

- [Go](https://golang.org/) - The programming language used.
- [Gofr](https://gofr.dev/) - The framework used.

## Authors

- Varuni Sood - Developer - ["varunisood06"](https://github.com/varunisood06)



