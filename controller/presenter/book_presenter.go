package presenter

import (
	"github.com/marutaku/epub-index-creator/domain"
	epubIndexCreator "github.com/marutaku/epub-index-creator/gen/epub_index_creator"
)

func ConvertBookToResponse(book *domain.Book) *epubIndexCreator.BookResponse {
	return &epubIndexCreator.BookResponse{
		Isbn:      epubIndexCreator.ISBN(book.ISBN),
		Title:     book.Title,
		Author:    book.Author,
		Language:  book.Language,
		Publisher: book.Publisher,
		Pages:     make([]*epubIndexCreator.PageResponse, 0), // TODO: Implement
	}
}
