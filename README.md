# RESTful-API-with-Go
this is beginner project to create API with Go and Gin Web Framework. This app provide sample API books.

## Prerequisite
- [Go language](https://golang.org/)
- [Gin framework](https://gin-gonic.com/docs/)

## How to install
- clone this repo
- run command below to install dependencies
```
go get .
```

## How to run
- inside project directory, run command below
```
go run main.go
```

## Endpoint
### `GET` `/books`
to get all books in database.

### `POST` `/books`
to insert new book

sample request body. `JSON`
```
{
    "id": "4",
    "title": "Harry Potter and the Philosophy Stone",
    "author": ["J.K. Rowling"],
    "price": 89.87
}
```

### `GET` `/books/:id`
to get book with specific id
