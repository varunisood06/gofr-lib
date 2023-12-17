package model

type Book struct {
    BookID   int    `json:"bookID"`
    Title    string `json:"title"`
    AuthorID int    `json:"authorID"`
    Issued   string `json:"issued"`
}
