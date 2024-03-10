package domain

import (
	"github.com/marutaku/epub-index-creator/ent"
)

type Book struct {
	ISBN      string
	Title     string
	Language  string
	Author    string
	Publisher string
	Pages     []*Page
}

func NewBook(isbn, title, language, author, publisher string, pages []*Page) *Book {
	return &Book{
		ISBN:      isbn,
		Title:     title,
		Language:  language,
		Author:    author,
		Publisher: publisher,
		Pages:     pages,
	}
}

func NewBookFromEnt(entBook *ent.Book) *Book {
	return NewBook(entBook.Isbn, entBook.Title, entBook.Language, entBook.Author, entBook.Publisher, nil)
}
