package database

import (
	"context"

	"github.com/marutaku/epub-index-creator/domain"
	"github.com/marutaku/epub-index-creator/ent"
	bookDB "github.com/marutaku/epub-index-creator/ent/book"
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

func (db *PageDatabase) Save(ctx context.Context, book *domain.Book, pageTitle, pagePath string) error {
	bookEntity, err := db.client.Book.Query().Where(bookDB.Isbn(book.ISBN)).Only(ctx)
	if err != nil {
		return err
	}
	_, err = db.client.Page.Create().
		SetBook(bookEntity).
		SetTitle(pageTitle).
		SetPath(pagePath).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
