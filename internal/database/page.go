package database

import (
	"context"

	"github.com/marutaku/epub-index-creator/indexer/ent"
	bookDB "github.com/marutaku/epub-index-creator/indexer/ent/book"
	"github.com/marutaku/epub-index-creator/indexer/internal/domain"
	_ "github.com/mattn/go-sqlite3"
)

type PageDatabase struct {
	client *ent.Client
}

func NewPageDatabase(sqliteDBPath string) *PageDatabase {
	client, err := ent.Open("sqlite3", sqliteDBPath)
	if err != nil {
		panic(err)
	}
	return &PageDatabase{client: client}
}

func (db *PageDatabase) Save(ctx context.Context, book *domain.Book, page *domain.Page) error {
	title, err := page.Title()
	if err != nil {
		return err
	}
	bookEntity, err := db.client.Book.Query().Where(bookDB.Isbn(book.ISBN)).Only(ctx)
	if err != nil {
		return err
	}
	_, err = db.client.Page.Create().
		SetBook(bookEntity).
		SetTitle(title).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
