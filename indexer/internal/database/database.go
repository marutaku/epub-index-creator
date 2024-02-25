package database

import "github.com/marutaku/epub-index-creator/indexer/internal/domain"

type Database interface {
	SaveBook(book domain.Book) error
}
