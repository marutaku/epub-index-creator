package database

import (
	"context"

	"github.com/marutaku/epub-index-creator/domain"
	"github.com/marutaku/epub-index-creator/ent"
	pageDB "github.com/marutaku/epub-index-creator/ent/page"
	_ "github.com/mattn/go-sqlite3"
)

type KeywordDatabase struct {
	client *ent.Client
}

func NewKeywordDatabase(sqliteDBPath string) *KeywordDatabase {
	client, err := ent.Open("sqlite3", sqliteDBPath)
	if err != nil {
		panic(err)
	}
	return &KeywordDatabase{client: client}
}

// Save saves a keyword to the database.
func (db *KeywordDatabase) Save(ctx context.Context, page *domain.Page, keyword string) error {
	title, err := page.Title()
	if err != nil {
		return err
	}
	pageEntity, err := db.client.Page.Query().Where(pageDB.Title(title)).Only(ctx)
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
