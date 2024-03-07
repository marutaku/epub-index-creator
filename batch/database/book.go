package database

import (
	"context"

	"github.com/marutaku/epub-index-creator/domain"
	"github.com/marutaku/epub-index-creator/ent"
	"github.com/marutaku/epub-index-creator/ent/book"
	_ "github.com/mattn/go-sqlite3"
)

type BookDatabase struct {
	client *ent.Client
}

func NewBookDatabase(sqliteDBPath string) *BookDatabase {
	client, err := ent.Open("sqlite3", sqliteDBPath)
	if err != nil {
		panic(err)
	}
	return &BookDatabase{client: client}
}

func (db *BookDatabase) FindBook(ctx context.Context, isbn string) (*domain.Book, error) {
	book, err := db.client.Book.Query().
		Where(book.Isbn(isbn)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return domain.NewBook(book.Isbn, book.Title, book.Language, book.Author, book.Publisher, nil), nil
}

func (db *BookDatabase) Save(ctx context.Context, book domain.Book) error {
	_, err := db.client.Book.Create().
		SetIsbn(book.ISBN).
		SetTitle(book.Title).
		SetLanguage(book.Language).
		SetAuthor(book.Author).
		SetPublisher(book.Publisher).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
