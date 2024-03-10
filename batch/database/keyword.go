package database

import (
	"context"

	"github.com/marutaku/epub-index-creator/domain"
	"github.com/marutaku/epub-index-creator/ent"
	pageDB "github.com/marutaku/epub-index-creator/ent/page"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	client *ent.Client
}

func NewKeywordDatabase(sqliteDBPath string) *Database {
	client, err := ent.Open("sqlite3", sqliteDBPath)
	if err != nil {
		panic(err)
	}
	return &Database{client: client}
}

// Save saves a keyword to the database.
func (db *Database) Save(ctx context.Context, page *domain.Page, keyword string) error {
	pageEntity, err := db.client.Page.Query().Where(pageDB.Title(page.Title)).Only(ctx)
	if err != nil {
		return err
	}
	_, err = db.client.Keyword.Create().
		SetPage(pageEntity).
		SetKeyword(keyword).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
