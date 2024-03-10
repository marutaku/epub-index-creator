package domain

import (
	"github.com/marutaku/epub-index-creator/ent"
)

type Page struct {
	Id    int
	Title string
	Path  string
}

func NewPage(id int, title, path string) *Page {
	return &Page{
		Id:    id,
		Title: title,
		Path:  path,
	}
}

func NewPageFromEnt(entPage *ent.Page) *Page {
	return NewPage(entPage.ID, entPage.Title, entPage.Path)
}
