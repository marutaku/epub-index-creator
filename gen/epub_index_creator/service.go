// Code generated by goa v3.15.0, DO NOT EDIT.
//
// epub_index_creator service
//
// Command:
// $ goa gen github.com/marutaku/epub-index-creator/design

package epubindexcreator

import (
	"context"
)

// Service is the epub_index_creator service interface.
type Service interface {
	// ListBooks implements ListBooks.
	ListBooks(context.Context, *ListBooksPayload) (res []*Book, err error)
	// FindBook implements FindBook.
	FindBook(context.Context, *FindBookPayload) (res *Book, err error)
	// CreateBook implements CreateBook.
	CreateBook(context.Context, *Book) (res *Book, err error)
	// UpdateBook implements UpdateBook.
	UpdateBook(context.Context, *UpdateBookPayload) (res *Book, err error)
	// DeleteBook implements DeleteBook.
	DeleteBook(context.Context, *DeleteBookPayload) (err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "epub-index-creator"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "epub_index_creator"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"ListBooks", "FindBook", "CreateBook", "UpdateBook", "DeleteBook"}

// Book is the result type of the epub_index_creator service FindBook method.
type Book struct {
	// ISBN of the book
	Isbn string
	// Title of the book
	Title string
	// Author of the book
	Author string
	// Language of the book
	Language string
	// Publisher of the book
	Publisher string
	// Pages of the book
	Pages []*Page
}

// DeleteBookPayload is the payload type of the epub_index_creator service
// DeleteBook method.
type DeleteBookPayload struct {
	// ISBN of the book
	Isbn string
	Book *Book
}

// FindBookPayload is the payload type of the epub_index_creator service
// FindBook method.
type FindBookPayload struct {
	// ISBN of the book
	Isbn string
}

type Keyword struct {
	// Keyword of the page
	Keyword string
}

// ListBooksPayload is the payload type of the epub_index_creator service
// ListBooks method.
type ListBooksPayload struct {
	// Maximum number of books to return
	Limit int
	// Field to paginate books
	Offset int
}

type Page struct {
	// Title of the page
	Title string
	// Keywords of the page
	Keywords []*Keyword
}

// UpdateBookPayload is the payload type of the epub_index_creator service
// UpdateBook method.
type UpdateBookPayload struct {
	// ISBN of the book
	Isbn string
	// Title of the book
	Title string
	// Author of the book
	Author string
	// Language of the book
	Language string
	// Publisher of the book
	Publisher string
}
