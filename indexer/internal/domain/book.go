package domain

import (
	"os"
	"path"

	"github.com/antchfx/xmlquery"
)

type Book struct {
	ISBN      string
	Title     string
	Language  string
	Author    string
	Publisher string
	Pages     []*Page
}

func NewBookFromOPF(filepath string) (*Book, error) {
	baseDir := path.Dir(filepath)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	doc, err := xmlquery.Parse(file)
	if err != nil {
		return nil, err
	}
	isbn := xmlquery.FindOne(doc, "//package/metadata/dc:identifier").InnerText()
	title := xmlquery.FindOne(doc, "//package/metadata/dc:title").InnerText()
	language := xmlquery.FindOne(doc, "//package/metadata/dc:language").InnerText()
	author := xmlquery.FindOne(doc, "//package/metadata/dc:creator").InnerText()
	publisher := xmlquery.FindOne(doc, "//package/metadata/dc:publisher").InnerText()
	items := xmlquery.Find(doc, "//package/manifest/item[@media-type='application/xhtml+xml']")
	pages := make([]*Page, 0, len(items))
	for _, item := range items {
		id := item.SelectAttr("id")
		pagePath := item.SelectAttr("href")
		pages = append(pages, NewPage(id, path.Join(baseDir, pagePath)))
	}
	return NewBook(isbn, title, language, author, publisher, pages), nil
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
