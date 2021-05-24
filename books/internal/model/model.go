package model

type Book struct {
	ID       string
	Title    string
	Pages    int
	AuthorID int
}

type BookAndAuthor struct {
	Book
	FirstName string
	LastName  string
}
